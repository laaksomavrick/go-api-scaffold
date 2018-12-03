package config

import (
	"os"
)

// Config defines the shape of the configuration values used across the api
type Config struct {
	Env        string
	Port       string
	DbUser     string
	DbPassword string
	DbName     string
}

// NewConfig returns the default configuration values used across the api
func NewConfig() *Config {
	// set defaults - these can be overwritten via command line

	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}

	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "postgres")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "postgres")
	}

	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "goals_development")
	}

	return &Config{
		Env:        os.Getenv("GO_ENV"),
		Port:       os.Getenv("PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
