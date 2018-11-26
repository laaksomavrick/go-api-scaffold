package main

import (
	"encoding/json"
	"net/http"

	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/user"
)

func main() {

	server := core.Server{
		Router: core.InitRouter(),
		DB:     core.InitDatabase(),
	}
	var routes = append(
		user.Routes,
	)
	server.Init(routes)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is up!")
}
