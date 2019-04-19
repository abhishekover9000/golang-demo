package driver

// Way to connect to the database

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

//ConnectDB connects to SQL Database
func ConnectDB() *sql.DB {
	var db *sql.DB

	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URI"))

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgURL)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	return db
}
