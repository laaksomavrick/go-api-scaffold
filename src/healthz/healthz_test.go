package healthz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laaksomavrick/goals-api/src/core"
)

func TestHealthzIndex(t *testing.T) {

	// TODO: confirm this runs middlewares / entire "app"

	server := &core.Server{
		Router: core.NewRouter(),
		DB:     core.NewDatabase(),
	}
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Errorf("GET /healthz error: %s", err)
	}

	rr := httptest.NewRecorder()

	handler := Index(server)
	handler.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("GET /healthz error: %d", res.StatusCode)
	}
}
