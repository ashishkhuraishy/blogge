package services

import (
	"fmt"
	"strconv"

	"github.com/ashishkhuraishy/blogge/src/domain/user"
	"github.com/ashishkhuraishy/blogge/src/utils/cryptoutils"
	"github.com/ashishkhuraishy/blogge/src/utils/datetime"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
)

var (
	// UserService is used to publically acces user service instance
	UserService userServiceInterface = &userService{}
)

// UserServiceInterface public interface for userService
type userServiceInterface interface {
	CreateUser(user.User) (*user.User, *resterror.RestError)
	GetUser(string) (*user.User, *resterror.RestError)
	UpdateUser(string, user.User, bool) (*user.User, *resterror.RestError)
	DeleteUser(string) *resterror.RestError
}

// UserService - to define user services
type userService struct{}

// CreateUser used to create a user
func (u *userService) CreateUser(user user.User) (*user.User, *resterror.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = cryptoutils.HashPassword(user.Password)
	user.DateCreated = datetime.GetCurrentFormattedTime()
	user.DateUpdated = user.DateCreated

	if restErr := user.Save(); restErr != nil {
		return nil, restErr
	}
	return &user, nil
}

// GetUser gets a specific user with the given id
// if any or else will return a not found error
func (u *userService) GetUser(idParam string) (*user.User, *resterror.RestError) {
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil || id < 1 {
		return nil, resterror.NewBadRequest(fmt.Sprintf("%s is not a valid user id", idParam))
	}

	usr := &user.User{
		ID: id,
	}

	restErr := usr.Get()
	return usr, restErr
}

// UpdateUser is used to update a user either fully or a single data about the
// the user. If the `isPartial` is true then only a part of the data will be
// updated where as if `isPartial` is false the whole new data will be added
// after removing all the old data
func (u *userService) UpdateUser(idParam string, usr user.User, isPartial bool) (*user.User, *resterror.RestError) {
	currentUser, restErr := u.GetUser(idParam)
	if restErr != nil {
		return nil, restErr
	}

	if usr.Password != "" {
		usr.Password = cryptoutils.HashPassword(usr.Password)
	}

	if isPartial {
		if usr.UserName == "" {
			usr.UserName = currentUser.UserName
		}
		if usr.Email == "" {
			usr.Email = currentUser.Email
		}
		if usr.Password == "" {
			usr.Password = currentUser.Password
		}
	}

	if err := usr.Validate(); err != nil {
		return nil, err
	}

	usr.ID = currentUser.ID
	usr.DateCreated = currentUser.DateCreated
	usr.DateUpdated = datetime.GetCurrentFormattedTime()

	restErr = usr.Update()
	return &usr, restErr
}

// DeleteUser will delete a user with the given id if the user exist
func (u *userService) DeleteUser(idParam string) *resterror.RestError {
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return resterror.NewBadRequest(fmt.Sprintf("%s is not a valid user id", idParam))
	}

	usr := &user.User{
		ID: id,
	}

	return usr.Delete()
}
