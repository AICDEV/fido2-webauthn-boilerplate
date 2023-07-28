package controllers

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/models"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/services"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/statics"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/storage"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
)

type MemberControllerInterface interface {
	GetRandomMemberName(ctx *gin.Context)
	AddAdditionalAuthenticatorBegin(ctx *gin.Context)
	AddAdditionalAuthenticatorFinish(ctx *gin.Context)
}

type MemberController struct {
	authenticationService services.AuthenticationServiceInterface
	fidoService           services.FidoServiceInterface
	storageDAOUser        storage.MongoDaoInterface
	members               []string
}

var (
	memberControllerSyncOnce sync.Once
	memberControllerInstance MemberControllerInterface
)

func GetMemberControllerInstance() MemberControllerInterface {
	memberControllerSyncOnce.Do(func() {
		memberControllerInstance = &MemberController{
			members: []string{
				"Hulk",
				"Capitan America",
				"Ironman",
				"Thor",
				"Loki",
				"Batman",
				"Superman",
			},
			authenticationService: services.GetAuthenticationServiceInstance(),
			fidoService:           services.GetFidoServiceInstance(),
			storageDAOUser: &storage.MongoDaoService{
				Database:       statics.GLOBAL_STATIC_DATABASE,
				Collection:     statics.GLOBAL_STATIC_USER_COLLECTION,
				StorageAdapter: storage.GetMongoStorageInstance(),
			},
		}
	})

	return memberControllerInstance
}

func (mc *MemberController) GetRandomMemberName(ctx *gin.Context) {
	member := models.MemberModel{
		Name: mc.getName(),
	}

	ctx.IndentedJSON(http.StatusOK, member)
}

func (mc *MemberController) AddAdditionalAuthenticatorBegin(ctx *gin.Context) {
	token, err := utils.FilterTokenStringFromAuthorizationHeader(ctx.Request)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("looks like your authorizatiion header is broken"))
		log.Println(err.Error())
		return
	}

	parsedUserClaims, err := mc.authenticationService.ParseClaimsFromToken(token)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("looks like your token is corrupt"))
		log.Println(err.Error())
		return
	}

	loaderUser := &models.User{}
	err = mc.storageDAOUser.QuerySingle("email", parsedUserClaims.Email, loaderUser)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("looks like your user data is broken"))
		log.Println(err.Error())
		return
	}

	options, err := mc.fidoService.AddAddionalAuthenticationDeviceBegin(loaderUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "can't register",
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, options)

}

func (mc *MemberController) AddAdditionalAuthenticatorFinish(ctx *gin.Context) {
	response, err := protocol.ParseCredentialCreationResponseBody(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation of payload failed",
		})

		return
	}

	err = mc.fidoService.AddAddionalAuthenticationDeviceFinish(response)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "adding of device failed",
		})

		return

	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"registration of device": "complete",
	})
}

func (mc *MemberController) getName() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(mc.members))

	return mc.members[index]
}
