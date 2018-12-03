package config

import (
	"os"
)

// Config defines the shape of the configuration values used across the api
type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	Port       string
}

// NewConfig returns the default configuration values used across the api
func NewConfig() *Config {
	// set defaults - these can be overwritten via command line
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "goals_development")
	os.Setenv("PORT", "3000")

	return &Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
	}
}
