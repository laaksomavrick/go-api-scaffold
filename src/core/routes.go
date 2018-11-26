package core

import (
	"net/http"
)

type ServerFunc func(s *Server) http.Handler

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc ServerFunc
}

type Routes []Route

// var routes = append(
// 	user.Routes,
// )

// inject routes from main
func routes(s *Server, routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(s)
		// handler = Logger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
