package user

import (
	"encoding/json"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

func Index(s *core.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from users!")
	})
}
