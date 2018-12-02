package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

// Create persists a user object to the database and returns the created record
func Create(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		err = user.validateForCreate()
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		err = user.prepareForInsert()
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		repo := newUserRepository(s.DB)
		err = repo.insert(&user)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
