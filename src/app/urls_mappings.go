package app

import "github.com/ashishkhuraishy/blogge/src/controller"

func mapUrls() {
	router.GET("/ping", controller.PingController.Ping)

	router.POST("/user", controller.UserController.Create)
	router.GET("/user/:user_id", controller.UserController.GetUser)
	router.PUT("/user/:user_id", controller.UserController.Update)
	router.PATCH("/user/:user_id", controller.UserController.Update)
	router.DELETE("/user/:user_id", controller.UserController.Delete)
}
