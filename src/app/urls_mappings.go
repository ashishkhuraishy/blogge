package app

import (
	"github.com/ashishkhuraishy/blogge/src/controller"
	"github.com/ashishkhuraishy/blogge/src/utils/middleware"
)

func mapUrls() {
	router.GET("/ping", controller.PingController.Ping)

	router.POST("/user", controller.UserController.Create)
	router.GET("/user/:user_id", middleware.AuthMiddleware(), controller.UserController.GetUser)
	router.PUT("/user/:user_id", controller.UserController.Update)
	router.PATCH("/user/:user_id", controller.UserController.Update)
	router.DELETE("/user/:user_id", controller.UserController.Delete)

}
