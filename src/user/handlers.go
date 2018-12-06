package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/util"
)

// Create persists a user object to the database and returns the created record
func Create(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		err = user.validateForCreate()
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		err = user.prepareForInsert()
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		repo := NewRepository(s.DB)
		err = repo.Insert(&user)
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

// Show parses the userId off the request context and sends back the full user object
func Show(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		userId := r.Context().Value("userId").(int)
		if userId == 0 {
			util.EncodeJSONError(w, errors.New("Error parsing request context for user"), http.StatusNotFound)
		}
		ur := NewRepository(s.DB)
		user, err := ur.FindById(userId)
		if err != nil {
			util.EncodeJSONError(w, errors.New("Error getting user"), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(user)
	}
}
