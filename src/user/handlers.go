package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

// Index returns an array of all users in JSON format
func Index(s *core.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from users!")
	})
}

// Create persists a user object to the database and returns the created record
func Create(s *core.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// todo logger
			core.JSONError(w, err, http.StatusBadRequest)
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			// todo logger
			core.JSONError(w, err, http.StatusBadRequest)
		}

		err = user.validateForCreate()
		if err != nil {
			// todo logger
			core.JSONError(w, err, http.StatusBadRequest)
			return
		}

		// todo service to hash pwd; repo for inserts

		query := `
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING id`

		err = s.DB.QueryRow(query, &user.Email, &user.Password).Scan(&user.ID)

		if err != nil {
			// todo logger
			core.JSONError(w, err, http.StatusBadRequest)
			panic(err)
		}

		json.NewEncoder(w).Encode(user)
	})
}
