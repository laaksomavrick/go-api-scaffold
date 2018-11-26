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

// InitDatabase verifies and returns a database connection
func InitDatabase() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("Unable to create the database object")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to the database")
	}
	fmt.Println("Connected to database")
	return db
}
