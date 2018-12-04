package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

		// sign and return jwt
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": dbUser.ID,
			"email":  dbUser.Email,
			"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// TODO: store this token in db; invalidate old token if relogging in
		// TODO: real hmac secret; read from file
		// if keyData, e := ioutil.ReadFile("test/hmacTestKey"); e == nil {
		// 	hmacSampleSecret = keyData
		// } else {
		// 	panic(e)
		// }
		tokenString, err := token.SignedString([]byte("asdqweasdqwe"))
		if err != nil {
			core.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})

	}
}
