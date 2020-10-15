package app

import "github.com/ashishkhuraishy/blogge/src/controller"

func mapUrls() {
	router.GET("/ping", controller.PingController.Ping)

	router.POST("/user", controller.UserController.Create)
	router.GET("/user/:user_id", controller.UserController.GetUser)
}
