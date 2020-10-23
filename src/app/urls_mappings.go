package app

import (
	"github.com/ashishkhuraishy/blogge/src/controller"
	"github.com/ashishkhuraishy/blogge/src/utils/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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
