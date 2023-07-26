package middleware

import (
	"net/http"
	"strings"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/services"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/statics"
	"github.com/gin-gonic/gin"
)

func LoadAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(statics.GLOBAL_AUTH_HEADER)

		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.Split(header, " ")

		if len(token) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO: validate token
		err := services.GetAuthenticationServiceInstance().VerifyToken(token[1])

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return

		}

		ctx.Next()

	}
}
