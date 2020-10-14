package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// PingController public instance for ping controller
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct{}

func (p *pingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
