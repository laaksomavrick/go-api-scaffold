package core

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Here to create the postgres db object
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "goals_development"
)

// NewDatabase verifies and returns a database connection
func NewDatabase() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Err creating db object: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Err pinging db: %s", err.Error())
	}
	return db
}
