package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/user"
)

func main() {

	// on server start
	// register routes and their middlewares;
	// register database connection;

	server := core.Server{
		Router: mux.NewRouter().StrictSlash(true),
	}
	var routes = append(
		user.Routes,
	)
	server.Init(routes)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is up!")
}
