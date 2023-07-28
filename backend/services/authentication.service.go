package services

import (
	"log"
	"sync"
	"time"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/models"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"github.com/golang-jwt/jwt/v4"
)

type AuthenticationServiceInterface interface {
	MakeToken(*models.User) (*models.TokenModel, error)
	ParseClaimsFromToken(string) (*ExtendedClaims, error)
	VerifyToken(string) error
}

type AuthenticationService struct {
	secret []byte
}

type CustomClaims struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"mail"`
}

type ExtendedClaims struct {
	jwt.RegisteredClaims
	CustomClaims
}

var (
	authenticationServiceSyncOnce sync.Once
	authenticationServiceInstance *AuthenticationService
)

func GetAuthenticationServiceInstance() AuthenticationServiceInterface {

	rawKeyPem, err := utils.LoadPrivateKeyFromDisk()

	if err != nil {
		log.Fatal((err.Error()))
	}

	parsedKey, err := jwt.ParseECPrivateKeyFromPEM(rawKeyPem)

	if err != nil {
		log.Fatal((err.Error()))
	}

	authenticationServiceSyncOnce.Do(func() {
		authenticationServiceInstance = &AuthenticationService{
			secret: parsedKey.X.Bytes(),
		}
	})

	return authenticationServiceInstance
}

func (as *AuthenticationService) MakeToken(user *models.User) (*models.TokenModel, error) {

	claims := ExtendedClaims{
		jwt.RegisteredClaims{
			Issuer:    "fido.workshop",
			Audience:  jwt.ClaimStrings{"fido.workshop"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Fido Workshop",
			ID:        user.ID,
		},
		CustomClaims{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(as.secret)

	if err != nil {
		return nil, err
	}

	return &models.TokenModel{
		Token: tokenString,
	}, nil
}

func (as *AuthenticationService) ParseClaimsFromToken(token string) (*ExtendedClaims, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &ExtendedClaims{}, func(t *jwt.Token) (interface{}, error) {
		return as.secret, nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken.Claims.(*ExtendedClaims), nil
}

func (as *AuthenticationService) VerifyToken(token string) error {
	_, err := jwt.ParseWithClaims(token, &ExtendedClaims{}, func(t *jwt.Token) (interface{}, error) {
		return as.secret, nil
	})

	return err
}
