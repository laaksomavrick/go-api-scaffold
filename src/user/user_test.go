package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laaksomavrick/goals-api/src/config"
	"github.com/laaksomavrick/goals-api/src/core"
)

func TestUserCreateWithValidPayload(t *testing.T) {

	config := config.NewConfig()

	server := core.NewServer(
		core.NewRouter(),
		core.NewDatabase(config),
		config,
	)

	_, err := server.DB.Exec("TRUNCATE TABLE users")
	if err != nil {
		panic(err)
	}

	server.Wire(Routes)

	body := map[string]string{
		"email":    "laakso.mavrick@gmail.com",
		"password": "qweqweqwe",
	}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(json))
	rr := httptest.NewRecorder()

	server.Router.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("POST /users error: %d", res.StatusCode)
	}

}

func TestUserCreateWithInvalidPayload(t *testing.T) {

	config := config.NewConfig()

	server := core.NewServer(
		core.NewRouter(),
		core.NewDatabase(config),
		config,
	)

	_, err := server.DB.Exec("TRUNCATE TABLE users")
	if err != nil {
		panic(err)
	}

	server.Wire(Routes)

	body := map[string]string{
		"email": "laakso.mavrick@gmail.com",
	}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(json))
	rr := httptest.NewRecorder()

	server.Router.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("POST /users error: %d", res.StatusCode)
	}

}

func TestUserCreateWithTooSmallPayload(t *testing.T) {

	config := config.NewConfig()

	server := core.NewServer(
		core.NewRouter(),
		core.NewDatabase(config),
		config,
	)

	_, err := server.DB.Exec("TRUNCATE TABLE users")
	if err != nil {
		panic(err)
	}

	server.Wire(Routes)

	body := map[string]string{
		"email":    "laakso.mavrick@gmail.com",
		"password": "less8",
	}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(json))
	rr := httptest.NewRecorder()

	server.Router.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("POST /users error: %d", res.StatusCode)
	}

}
