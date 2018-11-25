package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/laaksomavrick/goals-api/src/common"
)

func main() {

	// on server start
	// register routes and their middlewares;
	// register database connection;

	server := common.Server{
		Router: mux.NewRouter().StrictSlash(true),
	}
	server.Init()

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is up!")
}
