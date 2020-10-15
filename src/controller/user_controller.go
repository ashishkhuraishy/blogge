package controller

import (
	"net/http"

	"github.com/ashishkhuraishy/blogge/src/domain/user"
	"github.com/ashishkhuraishy/blogge/src/services"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
	"github.com/gin-gonic/gin"
)

var (
	// UserController : public instance to acces user controller
	UserController userControllerInterface = &userController{}
)

type userControllerInterface interface {
	Create(*gin.Context)
	GetUser(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userController struct{}

func (u *userController) Create(c *gin.Context) {
	var user user.User

	err := c.BindJSON(&user)
	if err != nil {
		resterr := resterror.NewBadRequest("invalid json")
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	result, resterr := services.UserService.CreateUser(user)
	if resterr != nil {
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUser converts takes a url parameter and checks if there is a valid user
// with that given id and returns the user
func (u *userController) GetUser(c *gin.Context) {
	idParam := c.Param("user_id")

	result, restErr := services.UserService.GetUser(idParam)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusFound, result)
}

// GetUser converts takes a url parameter and checks if there is a valid user
// with that given id and returns the user
func (u *userController) Update(c *gin.Context) {
	var user user.User
	idParam := c.Param("user_id")

	err := c.BindJSON(&user)
	if err != nil {
		resterr := resterror.NewBadRequest("invalid body")
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	result, restErr := services.UserService.UpdateUser(idParam, user, c.Request.Method == http.MethodPatch)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete converts takes a url parameter and checks if there is a valid user
// with that given id and returns the user
func (u *userController) Delete(c *gin.Context) {
	idParam := c.Param("user_id")

	restErr := services.UserService.DeleteUser(idParam)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "user removed sucessfully"})
}
