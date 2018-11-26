package healthz

import "github.com/laaksomavrick/goals-api/src/core"

// Routes defines the shape of all the routes for the healthz package
var Routes = core.Routes{
	core.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/healthz",
		HandlerFunc: Index,
	},
}
