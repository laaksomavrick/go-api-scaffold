package user

import (
	"encoding/json"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

// Index returns an array of all users in JSON format
func Index(s *core.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from users!")
	})
}
