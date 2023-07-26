package controllers

import (
	"net/http"
	"sync"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
)

type AuthenticationControllerInterface interface {
	Begin(ctx *gin.Context)
	Finish(ctx *gin.Context)
}

type AuthenticationController struct {
	authenticationService services.AuthenticationServiceInterface
	fidoService           services.FidoServiceInterface
}

var (
	authenticationControllerSyncOnce sync.Once
	authenticationControllerInstance AuthenticationControllerInterface
)

func GetAuthenticationControllerInstance() AuthenticationControllerInterface {
	authenticationControllerSyncOnce.Do(func() {
		authenticationControllerInstance = &AuthenticationController{
			fidoService:           services.GetFidoServiceInstance(),
			authenticationService: services.GetAuthenticationServiceInstance(),
		}
	})

	return authenticationControllerInstance
}

func (ac *AuthenticationController) Begin(ctx *gin.Context) {

	signInData := &models.ActionSignInRequest{}

	err := ctx.ShouldBindJSON(signInData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation of payload failed",
		})

		return
	}

	options, err := ac.fidoService.AuthenticateBegin(signInData)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't initiate auth dance",
		})
	}

	ctx.IndentedJSON(http.StatusOK, options)
}

func (ac *AuthenticationController) Finish(ctx *gin.Context) {
	response, err := protocol.ParseCredentialRequestResponseBody(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation of payload failed",
		})

		return
	}

	user, err := ac.fidoService.AuthenticateFinish(response)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "sign in failed",
		})

		return

	}

	token, err := ac.authenticationService.MakeToken(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"mesage": "crash",
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, token)

}
