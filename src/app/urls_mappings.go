package app

import (
	"github.com/ashishkhuraishy/blogge/src/controller"
	"github.com/ashishkhuraishy/blogge/src/utils/middleware"
)

func mapUrls() {
	router.GET("/ping", controller.PingController.Ping)

	router.POST("/signup", controller.UserController.Create)
	router.POST("/login", controller.UserController.Login)

	userGroup := router.Group("/user", middleware.AuthMiddleware())
	{
		userGroup.GET("/:user_id", controller.UserController.GetUser)
		userGroup.PUT("/:user_id", controller.UserController.Update)
		userGroup.PATCH("/:user_id", controller.UserController.Update)
		userGroup.DELETE("/:user_id", controller.UserController.Delete)
	}

}
