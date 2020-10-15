package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication will start the router and listens to traffic
func StartApplication() {
	mapUrls()

	router.Run(":8080")
}
