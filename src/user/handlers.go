package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

// Create persists a user object to the database and returns the created record
func Create(s *core.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			core.JSONError(w, err, http.StatusBadRequest)
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			core.JSONError(w, err, http.StatusBadRequest)
		}

		err = user.validateForCreate()
		if err != nil {
			core.JSONError(w, err, http.StatusBadRequest)
			return
		}

		user, err = user.prepareForInsert()
		if err != nil {
			core.JSONError(w, err, http.StatusBadRequest)
		}

		repo := newUserRepository(s.DB)
		err = repo.insert(&user)

		if err != nil {
			core.JSONError(w, err, http.StatusBadRequest)
			panic(err)
		}

		json.NewEncoder(w).Encode(user)
	})
}
