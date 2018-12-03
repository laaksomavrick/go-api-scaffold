package core

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/laaksomavrick/goals-api/src/config"
	_ "github.com/lib/pq" // Here to create the postgres db object
)

// NewDatabase verifies and returns a database connection
func NewDatabase(config *config.Config) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DbUser, config.DbPassword, config.DbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Err creating db object: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Err pinging db: %s", err.Error())
	}
	if config.Env != "testing" {
		fmt.Printf("Db: %s\t User: %s\t Password: %s\t\n", config.DbName, config.DbUser, config.DbPassword)
	}
	return db
}
