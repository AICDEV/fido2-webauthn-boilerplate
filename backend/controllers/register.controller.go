package controllers

import (
	"net/http"
	"sync"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
)

type RegisterControllerInterface interface {
	Begin(ctx *gin.Context)
	Finish(ctx *gin.Context)
}

type RegisterController struct {
	fidoService services.FidoServiceInterface
}

var (
	registerControllerSyncOnce sync.Once
	registerControllerInstance RegisterControllerInterface
)

func GetRegisterControllerInstance() RegisterControllerInterface {
	registerControllerSyncOnce.Do(func() {
		registerControllerInstance = &RegisterController{
			fidoService: services.GetFidoServiceInstance(),
		}
	})

	return registerControllerInstance
}

func (ac *RegisterController) Begin(ctx *gin.Context) {

	signUpData := &models.ActionSignUpRequest{}
	err := ctx.ShouldBindJSON(signUpData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation of payload failed",
		})

		return
	}

	options, err := ac.fidoService.SignUpBegin(signUpData)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't register",
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, options)
}

func (ac *RegisterController) Finish(ctx *gin.Context) {
	response, err := protocol.ParseCredentialCreationResponseBody(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "validation of payload failed",
		})

		return
	}

	err = ac.fidoService.SignUpFinish(response)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't process your request",
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"signup": "complete",
	})
}
