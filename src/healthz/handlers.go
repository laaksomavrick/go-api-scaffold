package healthz

import (
	"encoding/json"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
)

// Index returns the status of all the services for the api
func Index(s *core.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dbOk := true

		err := s.DB.Ping()
		if err != nil {
			dbOk = false
		}

		status := map[string]string{
			"server":   "ok",
			"postgres": boolToStatus(dbOk),
		}
		json.NewEncoder(w).Encode(status)
	}
}

func boolToStatus(b bool) string {
	if b == true {
		return "ok"
	}
	return "error"
}
