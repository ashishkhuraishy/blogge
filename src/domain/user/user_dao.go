package user

import "github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"

var (
	users = make(map[int64]*User, 0)
)

// Get : retrives a user with the id / returns a [RestError]
func (u *User) Get() *resterror.RestError {
	user := users[u.ID]
	if user == nil {
		return resterror.NewNotFoundError("user not found")
	}

	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password
	u.DateCreated = user.DateCreated
	u.DateUpdated = user.DateUpdated

	return nil
}

// Save used to save a user into the database
func (u *User) Save() *resterror.RestError {
	if users[u.ID] != nil {
		return resterror.NewBadRequest("user already exist")
	}

	users[u.ID] = u

	return nil
}
