package common

import (
	"net/http"

	"github.com/laaksomavrick/goals-api/src/user"
)

var routes = append(
	user.Routes,
)

func (s *Server) routes() {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)

		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
