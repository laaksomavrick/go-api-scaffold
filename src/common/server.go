package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server holds the shared dependencies of the service
type Server struct {
	// db     *someDatabase
	Router *mux.Router
}

// Init initializes the server instance
func (s *Server) Init() {
	s.routes()
	s.serve()
}

// Serve serves the application :)
func (s *Server) serve() {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(s.Router)))
}
