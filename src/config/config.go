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
	HmacSecret []byte
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

	if os.Getenv("HMAC_SECRET") == "" {
		// TODO: real hmac secret; read from file
		// TODO: should only need to read this once on app start, not every req
		// -> put hmacSecret in config
		// if keyData, e := ioutil.ReadFile("test/hmacTestKey"); e == nil {
		// 	hmacSampleSecret = keyData
		// } else {
		// 	panic(e)
		// }
		os.Setenv("HMAC_SECRET", "hmacsecret")
	}

	return &Config{
		Env:        os.Getenv("GO_ENV"),
		Port:       os.Getenv("PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		HmacSecret: []byte(os.Getenv("HMAC_SECRET")),
	}
}
