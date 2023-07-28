package app

import (
	"fmt"
	"log"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/controllers"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/middleware"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func StartApplication() {
	log.Println("booting backend")

	envConfig := utils.ParseEnv()

	router := gin.New()
	api := router.Group("/api")

	apiV1 := api.Group("/v1")
	apiV1.Use(gin.Logger())

	serviceApi := apiV1.Group("/service")
	memberApi := apiV1.Group("/member")

	memberApi.Use(middleware.LoadAuthMiddleware())

	serviceApi.POST("/signup/begin", controllers.GetRegisterControllerInstance().Begin)
	serviceApi.POST("/signup/finish", controllers.GetRegisterControllerInstance().Finish)

	serviceApi.POST("/authenticate/begin", controllers.GetAuthenticationControllerInstance().Begin)
	serviceApi.POST("/authenticate/finish", controllers.GetAuthenticationControllerInstance().Finish)

	memberApi.GET("/name", controllers.GetMemberControllerInstance().GetRandomMemberName)
	memberApi.POST("/register/device/begin", controllers.GetMemberControllerInstance().AddAdditionalAuthenticatorBegin)
	memberApi.POST("/register/device/finish", controllers.GetMemberControllerInstance().AddAdditionalAuthenticatorFinish)

	log.Println("backend setup complete. try serving now...")
	log.Printf("PORT: %d, HOST: %s", envConfig.Port, envConfig.Host)

	err := router.Run(fmt.Sprintf("%v:%v", envConfig.Host, envConfig.Port))
	if err != nil {
		log.Fatal(err.Error())
	}

}
