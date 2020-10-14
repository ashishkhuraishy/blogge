package services

import "github.com/ashishkhuraishy/blogge/domain/user"

var (
	// UserService is used to publically acces user service instance
	UserService userServiceInterface = &userService{}
)

// UserServiceInterface public interface for userService
type userServiceInterface interface {
	CreateUser(user.User)
}

// UserService - to define user services
type userService struct{}

// CreateUser used to create a user
func (u *userService) CreateUser(user user.User) {
	user.Save()
}
