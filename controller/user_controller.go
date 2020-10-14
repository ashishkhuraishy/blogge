package controller

import (
	"net/http"

	"github.com/ashishkhuraishy/blogge/domain/user"
	"github.com/gin-gonic/gin"
)

var (
	// UserController : public instance to acces user controller
	UserController userControllerInterface = &userController{}
)

type userControllerInterface interface {
	Create(*gin.Context)
}

type userController struct{}

func (u *userController) Create(c *gin.Context) {
	var user user.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Json"})
		return
	}
}
