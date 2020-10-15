package services

import (
	"github.com/ashishkhuraishy/blogge/src/domain/user"
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
	GetUser(int64) (*user.User, *resterror.RestError)
}

// UserService - to define user services
type userService struct{}

// CreateUser used to create a user
func (u *userService) CreateUser(user user.User) (*user.User, *resterror.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = datetime.GetCurrentFormattedTime()
	user.DateUpdated = user.DateCreated

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser gets a specific user with the given id
// if any or else will return a not found error
func (u *userService) GetUser(id int64) (*user.User, *resterror.RestError) {
	usr := &user.User{
		ID: id,
	}

	if err := usr.Get(); err != nil {
		return nil, err
	}

	return usr, nil
}
