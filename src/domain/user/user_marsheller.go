package user

import (
	"encoding/json"

	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
)

// PrivateUser contains all info about the user
// only avaibale to the owner user
type PrivateUser struct {
	ID          int64  `json:"id"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

// PublicUser contains all the info about a
// user excluding some private/sensitive info
type PublicUser struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
}

// Marshaller will marshall the incoming user
// and convert it into private/public user
func (u *User) Marshaller(isPrivate bool) interface{} {
	usrJSON, err := json.Marshal(u)
	if err != nil {
		return resterror.NewInternalServerError("could not marshall the data")
	}

	if isPrivate {
		var user PrivateUser
		json.Unmarshal(usrJSON, &user)
		return user
	}

	var user PublicUser
	json.Unmarshal(usrJSON, &user)
	return user
}
