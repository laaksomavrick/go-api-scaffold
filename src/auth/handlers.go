package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/user"
)

// Create authenticates a user and returns a jwt for subsequent requests
func Create(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		// marshal user
		var u user.User
		err = json.Unmarshal(body, &u)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		// find user for email
		ur := user.NewRepository(s.DB)
		dbUser, err := ur.FindByEmail(u.Email)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusNotFound)
			return
		}

		// verify input password to hashed password
		err = dbUser.CompareHashAndPassword(u.Password)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusUnauthorized)
			return
		}

		// generate token
		service := NewService()
		token, err := service.GenerateToken(dbUser)
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})

	}
}
