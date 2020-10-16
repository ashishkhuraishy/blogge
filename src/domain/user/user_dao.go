package user

import (
	"log"

	"github.com/ashishkhuraishy/blogge/src/databases/psql"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/psqlerrors"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
)

var (
	users = make(map[int64]*User, 0)
)

const (
	queryInsertUser = `INSERT INTO users(username, email, password, date_created, date_updated) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	queryGetUser    = `SELECT id, username, email, password, date_created, date_updated FROM users WHERE id=$1;`
	queryUpdateUser = `UPDATE users SET username=$1, email=$2, password=$3, date_created=$4, date_updated=$5 WHERE id=$6`
	queryDeleteUser = `DELETE FROM users WHERE id=$1`
)

// Get : retrives a user with the id / returns a [RestError]
func (u *User) Get() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryGetUser)
	defer stmnt.Close()
	if err != nil {
		log.Println(err)
		return psqlerrors.ParseErrors(err)
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
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Save used to save a user into the database
func (u *User) Save() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryInsertUser)
	defer stmnt.Close()
	if err != nil {
		log.Println(err)
		return psqlerrors.ParseErrors(err)
	}

	err = stmnt.QueryRow(
		u.UserName,
		u.Email,
		u.Password,
		u.DateCreated,
		u.DateUpdated,
	).Scan(&u.ID)

	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Update updates the user with the given info
func (u *User) Update() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryUpdateUser)
	defer stmnt.Close()
	if err != nil {
		return psqlerrors.ParseErrors(err)
	}

	_, err = stmnt.Exec(
		u.UserName,
		u.Email,
		u.Password,
		u.DateCreated,
		u.DateUpdated,
		u.ID,
	)

	if err != nil {
		log.Println(err)
		return psqlerrors.ParseErrors(err)
	}

	return nil
}

// Delete used to delete a user from the db
func (u *User) Delete() *resterror.RestError {
	stmnt, err := psql.Client.Prepare(queryDeleteUser)
	defer stmnt.Close()
	if err != nil {
		log.Println(err)
		return psqlerrors.ParseErrors(err)
	}

	_, err = stmnt.Exec(u.ID)
	if err != nil {
		log.Println(err)
		return psqlerrors.ParseErrors(err)
	}

	return nil
}
