package core

import (
	"net/http"
)

// ServerFunc defines the shape of handler fns on routes.
// Server is injected for common access to routes/db/logger/etc
type ServerFunc func(s *Server) http.Handler

// Route defines the shape of a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc ServerFunc
}

// Routes defines the shape of an array of routes
type Routes []Route

func routes(s *Server, routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(s)
		// todo request logger
		// handler = Logger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
