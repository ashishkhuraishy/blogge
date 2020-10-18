package app

import (
	"github.com/ashishkhuraishy/blogge/src/controller"
	"github.com/ashishkhuraishy/blogge/src/utils/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Blogge API Docs
// @version 1.0
// @description API Docs for Golang Project Blogge.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email ashishkhuraishy@gmail.com

// @license.name MIT
// @license.url https://github.com/ashishkhuraishy/blogge/blob/master/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func mapUrls() {
	router.GET("/ping", controller.PingController.Ping)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/signup", controller.UserController.SignUp)
	router.POST("/login", controller.UserController.Login)

	userGroup := router.Group("/user", middleware.AuthMiddleware())
	{
		userGroup.GET("/:user_id", controller.UserController.GetUser)
		userGroup.PUT("/:user_id", controller.UserController.Update)
		userGroup.PATCH("/:user_id", controller.UserController.Update)
		userGroup.DELETE("/:user_id", controller.UserController.Delete)
	}

}
