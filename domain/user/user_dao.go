package user

import "fmt"

var (
	users = make(map[int64]*User, 0)
)

// Save used to save a user into the database
func (u *User) Save() {
	users[u.ID] = u
	fmt.Println(users[u.ID])
}
