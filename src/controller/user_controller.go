package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
	SignUp(*gin.Context)
	Login(*gin.Context)
	GetUser(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userController struct{}

func (u *userController) SignUp(c *gin.Context) {
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
	id, _, isOwner, restErr := extractUserInfo(c)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, restErr := services.UserService.GetUser(id)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusFound, result.Marshaller(isOwner))
}

// GetUser converts takes a url parameter and checks if there is a valid user
// with that given id and returns the user
func (u *userController) Update(c *gin.Context) {
	var user user.User
	id, _, isOwner, restErr := extractUserInfo(c)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	if !isOwner {
		restErr = resterror.NewInvalidCredentialsError()
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	err := c.BindJSON(&user)
	if err != nil {
		resterr := resterror.NewBadRequest("invalid body")
		c.JSON(resterr.StatusCode, resterr)
		return
	}

	result, restErr := services.UserService.UpdateUser(id, user, c.Request.Method == http.MethodPatch)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshaller(true))
}

// Delete converts takes a url parameter and checks if there is a valid user
// with that given id and returns the user
func (u *userController) Delete(c *gin.Context) {
	id, _, isOwner, restErr := extractUserInfo(c)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	if !isOwner {
		restErr = resterror.NewInvalidCredentialsError()
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	restErr = services.UserService.DeleteUser(id)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "user removed sucessfully"})
}

// -------------------- Helper Methods ---------------------- //

// GetDataFromMiddleware is a helper fn used to extract data
// from the auth middleware
func getDataFromMiddleWare(c *gin.Context) (int64, bool, *resterror.RestError) {
	requestedUser := int64(c.GetFloat64("user_id"))
	isAdmin := c.GetBool("is_admin")

	if requestedUser == 0 {
		return 0, false, resterror.NewInvalidCredentialsError()
	}

	return requestedUser, isAdmin, nil
}

// ExtractUser id will extract a user id from the given context
func extractUserID(userParam string) (int64, *resterror.RestError) {
	id, err := strconv.ParseInt(userParam, 10, 64)

	if err != nil || id < 1 {
		return 0, resterror.NewBadRequest(fmt.Sprintf("%s is not a valid user id", userParam))
	}

	return id, nil
}

// ExtractuserInfo will take in the current context then will return
// the `userid` at the request, the `id` of the requested user from the
// token and a bool field to check the request is from the owner & finally
// if anthing goes wrong an error interface
func extractUserInfo(c *gin.Context) (int64, int64, bool, *resterror.RestError) {
	idParam := c.Param("user_id")

	id, restErr := extractUserID(idParam)
	if restErr != nil {
		return 0, 0, false, restErr
	}

	requestedUser, _, restErr := getDataFromMiddleWare(c)
	if restErr != nil {
		return 0, 0, false, restErr
	}

	return id, requestedUser, id == requestedUser, nil
}
