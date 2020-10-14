package domain

import "fmt"

var (
	users = make(map[int64]User, 0)
)

// Save used to save a user into the database
func (u *User) Save() {
	fmt.Println(users[u.ID])
}
