package user

import (
	"strings"

	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
)

// User defines the core User of the app
type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

// AuthResponse contains the user data that will be returned
// after signup / login
type AuthResponse struct {
	Profile interface{} `json:"profile"`
	Token   string      `json:"token"`
}

// Validate is used to validate wheather
// the current user provided is a valid
// user or not
func (u *User) Validate() *resterror.RestError {
	u.Email = strings.TrimSpace(u.Email)
	u.UserName = strings.TrimSpace(u.UserName)

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if u.UserName == "" {
		return resterror.NewBadRequest("username cannot be empty")
	}
	if len(u.Password) < 6 {
		return resterror.NewBadRequest("password is too short")
	}

	return nil
}

// ValidateEmail will trim and check if the email is valid
// or not and returns a rest error if invalid
func (u *User) ValidateEmail() *resterror.RestError {
	u.Email = strings.ToLower(u.Email)
	u.Email = strings.TrimSpace(u.Email)

	if u.Email == "" || !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
		return resterror.NewBadRequest("invalid email")
	}

	return nil
}
