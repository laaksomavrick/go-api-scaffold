package healthz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laaksomavrick/goals-api/src/config"
	"github.com/laaksomavrick/goals-api/src/core"
)

func TestHealthzIndex(t *testing.T) {

	// initializes env
	config := config.NewConfig()

	// create a new server instance
	server := core.NewServer(
		core.NewRouter(),
		core.NewDatabase(config),
		config,
	)
	// add the healthz routes and apply middlewares
	server.Wire(Routes)

	// create a request
	req, _ := http.NewRequest("GET", "/healthz", nil)
	// create a request recorder
	rr := httptest.NewRecorder()

	// serve the healthz routes
	server.Router.ServeHTTP(rr, req)

	// get the result
	res := rr.Result()

	// make an assertion
	if res.StatusCode != http.StatusOK {
		t.Errorf("GET /healthz error: %d", res.StatusCode)
	}
}
