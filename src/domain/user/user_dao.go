package user

import (
	"log"

	"github.com/ashishkhuraishy/blogge/src/databases/psql"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
)

var (
	users = make(map[int64]*User, 0)
)

const (
	queryInsertUser    = `INSERT INTO users(username, email, password, date_created, date_updated) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	queryGetUserWithID = `SELECT id, username, email, password, date_created, date_updated FROM users WHERE id=$1;`
)

// Get : retrives a user with the id / returns a [RestError]
func (u *User) Get() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryGetUserWithID)
	defer stmnt.Close()
	if err != nil {
		log.Println(err)
		return resterror.NewInternalServerError("database_error")
	}

	err = stmnt.QueryRow(
		u.ID,
	).Scan(
		&u.ID,
		&u.UserName,
		&u.Email,
		&u.Password,
		&u.DateCreated,
		&u.DateUpdated,
	)

	if err != nil {
		log.Println(err)
		return resterror.NewInternalServerError("database_error")
	}

	return nil
}

// Save used to save a user into the database
func (u *User) Save() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryInsertUser)
	defer stmnt.Close()
	if err != nil {
		log.Println(err)
		return resterror.NewInternalServerError("database_error")
	}

	err = stmnt.QueryRow(
		u.UserName,
		u.Email,
		u.Password,
		u.DateCreated,
		u.DateUpdated,
	).Scan(&u.ID)

	if err != nil {
		return resterror.NewInternalServerError("database_error")
	}

	return nil
}

// Delete used to delete a user from the db
func (u *User) Delete() *resterror.RestError {

	return nil
}
