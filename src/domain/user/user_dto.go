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

// Validate is used to validate wheather
// the current user provided is a valid
// user or not
func (u *User) Validate() *resterror.RestError {
	u.Email = strings.TrimSpace(u.Email)
	u.UserName = strings.TrimSpace(u.UserName)

	if u.Email == "" {
		return resterror.NewBadRequest("email cannot be empty")
	}
	if u.UserName == "" {
		return resterror.NewBadRequest("username cannot be empty")
	}
	if len(u.Password) < 6 {
		return resterror.NewBadRequest("password is too short")
	}

	return nil
}
