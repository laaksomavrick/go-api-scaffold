package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/laaksomavrick/goals-api/src/util"
)

// CheckAuthentication parses the Authorization header and populates the request context with userId
func CheckAuthentication(next http.Handler, authRequired bool, hmacSecret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check if route is guarded by require auth
		if authRequired == false {
			next.ServeHTTP(w, r)
			return
		}

		// get token
		bearerTokenString := r.Header.Get("Authorization")
		if bearerTokenString == "" {
			util.EncodeJSONError(w, errors.New("Invalid Authorization Header"), http.StatusBadRequest)
			return
		}
		split := strings.Split(bearerTokenString, "Bearer ")
		if len(split) < 2 {
			util.EncodeJSONError(w, errors.New("Invalid Authorization Header"), http.StatusBadRequest)
			return
		}
		token := split[1]

		// parse token
		user, err := parseToken(token, hmacSecret)
		if err != nil {
			util.EncodeJSONError(w, err, http.StatusInternalServerError)
			return
		}

		// attach user obj to request for req.user
		ctx := context.WithValue(r.Context(), "userId", user)

		// next
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func parseToken(tokenString string, hmacSecret []byte) (int, error) {
	var userID int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return userID, errors.New("Error parsing token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = int(claims["userId"].(float64))
		if err != nil {
			return userID, err
		}
		return userID, nil
	} else {
		return userID, errors.New("Error parsing token")
	}
}
