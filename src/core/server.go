package core

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server holds the essential shared dependencies of the service
type Server struct {
	Router *mux.Router
	DB     *sql.DB
	// todo logger obj for debug / misc server logs?
}

// NewServer constructs a new instance of a server
func NewServer(router *mux.Router, db *sql.DB) *Server {
	return &Server{
		Router: router,
		DB:     db,
	}
}

// Init applies the middleware stack, registers route handlers, and serves the application
func (s *Server) Init(routes Routes) {
	s.Wire(routes)
	s.Serve()
}

// Serve serves the application :)
func (s *Server) Serve() {
	log.Fatal(http.ListenAndServe(":3000", s.Router))
}

// Wire applies middlewares to all routes and registers them to the Server.Router
func (s *Server) Wire(routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(s)
		handler = StdLogger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		headersOk := handlers.AllowedHeaders([]string{"Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		handlers.CORS(originsOk, headersOk, methodsOk)(s.Router)

	}
}
