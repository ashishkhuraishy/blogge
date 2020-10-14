package user

import (
	"strings"

	"github.com/ashishkhuraishy/blogge/utils/errors/resterror"
)

// User defines the core User of the app
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

// Validate is used to validate wheather
// the current user provided is a valid user or
// not
func (u *User) Validate() *resterror.RestError {
	u.Email = strings.TrimSpace(u.Email)
	u.Name = strings.TrimSpace(u.Name)

	if u.ID < 1 {
		return resterror.NewBadRequest("invalid id")
	}
	if u.Email == "" {
		return resterror.NewBadRequest("email cannot be empty")
	}
	if u.Name == "" {
		return resterror.NewBadRequest("name cannot be empty")
	}
	if len(u.Password) < 6 {
		return resterror.NewBadRequest("password is too short")
	}

	return nil
}
