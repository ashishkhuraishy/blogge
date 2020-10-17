package psql

import (
	"log"
)

const (
	// Query to create a table if does not exist
	queryCreateUserTable = `CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(100) NOT NULL,
		date_created TIMESTAMP NOT NULL,
		date_updated TIMESTAMP
	);`
)

func schemaInit() error {
	_, err := Client.Exec(queryCreateUserTable)
	if err != nil {
		log.Println("Here " + err.Error())
		return err
	}

	return nil
}
