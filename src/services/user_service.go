package services

import (
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
	LoginService(user.User) (*user.User, *resterror.RestError)
	GetUser(int64) (*user.User, *resterror.RestError)
	UpdateUser(int64, user.User, bool) (*user.User, *resterror.RestError)
	DeleteUser(int64) *resterror.RestError
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
func (u *userService) GetUser(id int64) (*user.User, *resterror.RestError) {

	usr := &user.User{
		ID: id,
	}

	restErr := usr.Get()
	return usr, restErr
}

// LoginService recives a user with username and pass and will try to authenticate
// the user with the data on db
func (u *userService) LoginService(usr user.User) (*user.User, *resterror.RestError) {
	if err := usr.ValidateEmail(); err != nil {
		return nil, err
	}

	if usr.Password == "" {
		return nil, resterror.NewInvalidCredentialsError()
	}

	currentUser, err := usr.GetUserByEmail()
	if err != nil {
		return nil, resterror.NewInvalidCredentialsError()
	}

	if !cryptoutils.VerifyPasswordAndHash(usr.Password, currentUser.Password) {
		return nil, resterror.NewInvalidCredentialsError()
	}

	return currentUser, nil
}

// UpdateUser is used to update a user either fully or a single data about the
// the user. If the `isPartial` is true then only a part of the data will be
// updated where as if `isPartial` is false the whole new data will be added
// after removing all the old data
func (u *userService) UpdateUser(id int64, usr user.User, isPartial bool) (*user.User, *resterror.RestError) {
	currentUser, restErr := u.GetUser(id)
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
func (u *userService) DeleteUser(id int64) *resterror.RestError {
	usr := &user.User{
		ID: id,
	}

	return usr.Delete()
}
