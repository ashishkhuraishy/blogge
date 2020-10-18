package controller

import (
	"errors"
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
	Login(*gin.Context)
}

type userController struct{}

func (u *userController) Create(c *gin.Context) {
	var usr user.User

	err := c.BindJSON(&usr)
	if err != nil {
		resterr := resterror.NewBadRequest("invalid json")
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	result, resterr := services.UserService.CreateUser(usr)
	if resterr != nil {
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	authService := services.JWTAuthService()
	token := authService.GenerateToken(result.ID, false)

	c.JSON(http.StatusOK, gin.H{"profile": result.Marshaller(true), "token": token})
}

// Login used to validate a user and generate a token
func (u *userController) Login(c *gin.Context) {
	var usr user.User

	err := c.BindJSON(&usr)
	if err != nil {
		resterr := resterror.NewBadRequest("invalid json")
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	result, resterr := services.UserService.LoginService(usr)
	if resterr != nil {
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	authService := services.JWTAuthService()
	token := authService.GenerateToken(result.ID, false)

	c.JSON(http.StatusOK, gin.H{"profile": result.Marshaller(true), "token": token})
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

// GetDataFromMiddleware is a helper fn used to extract data
// from the auth middleware
func getDataFromMiddleWare(c *gin.Context) (int64, bool, error) {
	requestedUser := int64(c.GetFloat64("user_id"))
	isAdmin := c.GetBool("is_admin")

	if requestedUser == 0 {
		return 0, false, errors.New("corrupted token")
	}

	return requestedUser, isAdmin, nil
}
