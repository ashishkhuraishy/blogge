package psql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgress diver
)

// Env Variable Names
const (
	psqlUserName = "PSQL_USER_NAME"
	psqlPassword = "PSQL_PASSWORD"
	psqlPort     = "PSQL_PORT"
	psqlDBName   = "PSQL_DB_NAME"
	psqlHost     = "PSQL_DB_HOST"
)

var (
	// Client used a public variable to acces the client db
	Client *sql.DB
)

func init() {
	var err error

	// Load env variables if any locally
	godotenv.Load()

	host := os.Getenv(psqlHost)
	if host == "" {
		host = "localhost"
	}

	// Generating the connection string
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", os.Getenv(psqlUserName), os.Getenv(psqlPassword), os.Getenv(psqlDBName), host)
	log.Println(connectionStr)
	Client, err = sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	err = schemaInit()
	if err != nil {
		panic(err)
	}

	log.Println("Database Started succesfully")
}
