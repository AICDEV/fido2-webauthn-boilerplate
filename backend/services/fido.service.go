package services

import (
	"log"
	"sync"
	"time"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/statics"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/storage"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/utils"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"go.mongodb.org/mongo-driver/bson"
)

type FidoServiceInterface interface {
	SignUpBegin(*models.ActionSignUpRequest) (*protocol.CredentialCreation, error)
	SignUpFinish(*protocol.ParsedCredentialCreationData) error
	AuthenticateBegin(*models.ActionSignInRequest) (*protocol.CredentialAssertion, error)
	AuthenticateFinish(*protocol.ParsedCredentialAssertionData) (*models.User, error)
	AddAddionalAuthenticationDeviceBegin(*models.User) (*protocol.CredentialCreation, error)
	AddAddionalAuthenticationDeviceFinish(*protocol.ParsedCredentialCreationData) error
}

type FidoService struct {
	webAuthInst           *webauthn.WebAuthn
	storageCacheChallenge storage.RedisStorageInterface
	storageDAOUser        storage.MongoDaoInterface
}

var (
	fidoServiceSyncOnce sync.Once
	fidoServiceInstance *FidoService
)

func GetFidoServiceInstance() FidoServiceInterface {
	fidoServiceSyncOnce.Do(func() {
		wInst, err := webauthn.New(getWebAuthConfig())
		if err != nil {
			log.Fatal(err.Error())
		}

		fidoServiceInstance = &FidoService{
			webAuthInst: wInst,
			storageDAOUser: &storage.MongoDaoService{
				Database:       statics.GLOBAL_STATIC_DATABASE,
				Collection:     statics.GLOBAL_STATIC_USER_COLLECTION,
				StorageAdapter: storage.GetMongoStorageInstance(),
			},
			storageCacheChallenge: storage.GetRedisClientInstance(),
		}
	})

	return fidoServiceInstance
}

func (fs *FidoService) SignUpBegin(signUpData *models.ActionSignUpRequest) (*protocol.CredentialCreation, error) {
	existingUser := &models.User{}
	err := fs.storageDAOUser.QuerySingle(
		"email",
		utils.FilterStorageQuery(signUpData.Email),
		existingUser,
	)

	if err == nil {
		return nil, err
	}

	user := fs.newUser(
		signUpData.Email,
		signUpData.DisplayName,
		signUpData.Firstname,
		signUpData.Lastname,
	)

	options, _, err := fs.createRegistrationChallengeAndCache(user)

	if err != nil {
		return nil, err
	}

	return options, nil

}

func (fs *FidoService) SignUpFinish(response *protocol.ParsedCredentialCreationData) error {

	cache := &models.CacheModel{}

	data, err := fs.storageCacheChallenge.Get(response.Response.CollectedClientData.Challenge)
	if err != nil {
		return err
	}

	err = utils.DecodeFromRedisStore(data, cache)

	if err != nil {
		return err
	}

	credential, err := fs.webAuthInst.CreateCredential(&cache.User, *cache.SessionData, response)

	if err != nil {
		return err
	}

	cache.User.Credentials = append(cache.User.Credentials, *credential)
	cache.User.RegistrationID = []byte{0x00}
	cache.SessionData = nil

	fs.storageDAOUser.Save(
		cache.User,
	)

	return nil
}

func (fs *FidoService) AuthenticateBegin(signInData *models.ActionSignInRequest) (*protocol.CredentialAssertion, error) {
	user := &models.User{}

	err := fs.storageDAOUser.QuerySingle(
		"email",
		utils.FilterStorageQuery(signInData.Email),
		user,
	)

	if err != nil {
		return nil, err
	}

	options, session, err := fs.webAuthInst.BeginLogin(user)

	if err != nil {
		return nil, err
	}

	cache := &models.CacheModel{
		User:        *user,
		SessionData: session,
	}

	enc, err := utils.EncodeForRedisStore(cache)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = fs.storageCacheChallenge.SetWithTTL(session.Challenge, string(enc), time.Second*30)

	if err != nil {
		log.Fatal(err.Error())
	}

	return options, nil

}
func (fs *FidoService) AuthenticateFinish(response *protocol.ParsedCredentialAssertionData) (*models.User, error) {

	cache := &models.CacheModel{}

	data, err := fs.storageCacheChallenge.Get(response.Response.CollectedClientData.Challenge)
	if err != nil {
		return nil, err
	}

	err = utils.DecodeFromRedisStore(data, cache)

	if err != nil {
		return nil, err
	}

	_, err = fs.webAuthInst.ValidateLogin(&cache.User, *cache.SessionData, response)

	if err != nil {
		return nil, err
	}

	return &cache.User, nil
}

func (fs *FidoService) AddAddionalAuthenticationDeviceBegin(user *models.User) (*protocol.CredentialCreation, error) {

	options, _, err := fs.createRegistrationChallengeAndCache(user)

	if err != nil {
		return nil, err
	}

	return options, nil
}

func (fs *FidoService) AddAddionalAuthenticationDeviceFinish(response *protocol.ParsedCredentialCreationData) error {
	cache := &models.CacheModel{}

	data, err := fs.storageCacheChallenge.Get(response.Response.CollectedClientData.Challenge)
	if err != nil {
		return err
	}

	err = utils.DecodeFromRedisStore(data, cache)

	if err != nil {
		return err
	}

	credential, err := fs.webAuthInst.CreateCredential(&cache.User, *cache.SessionData, response)

	if err != nil {
		return err
	}

	cache.User.Credentials = append(cache.User.Credentials, *credential)
	cache.User.RegistrationID = []byte{0x00}
	cache.SessionData = nil

	updatedUser := &models.User{}

	err = fs.storageDAOUser.UpdateById(
		cache.User.ID,
		updatedUser,
		bson.M{
			"credentials": cache.User.Credentials,
		},
	)

	if err != nil {
		return err
	}

	return nil

}

func (fs *FidoService) createRegistrationChallengeAndCache(user *models.User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	options, session, err := fs.webAuthInst.BeginRegistration(user)

	cache := &models.CacheModel{
		User:        *user,
		SessionData: session,
	}

	enc, err := utils.EncodeForRedisStore(cache)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = fs.storageCacheChallenge.SetWithTTL(session.Challenge, string(enc), time.Second*30)

	if err != nil {
		log.Fatal(err.Error())
	}

	return options, session, err
}

func (fs *FidoService) newUser(email string, displayName string, firstName string, lastName string) *models.User {
	return GetUserServiceInstance().New(
		email,
		displayName,
		firstName,
		lastName,
	)
}

func getWebAuthConfig() *webauthn.Config {
	return &webauthn.Config{
		RPDisplayName: "Fido Workshop",
		RPID:          "fido.workshop",
		RPOrigins: []string{
			"https://fido.workshop",
		},
		Timeouts: webauthn.TimeoutsConfig{
			Login: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 45,
				TimeoutUVD: time.Second * 45,
			},
			Registration: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 45,
				TimeoutUVD: time.Second * 45,
			},
		},
	}

}
