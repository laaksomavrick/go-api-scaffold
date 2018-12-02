package user

import "github.com/laaksomavrick/goals-api/src/core"

// Routes defines the shape of all the routes for the user package
var Routes = core.Routes{
	core.Route{
		Name:        "Create",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: Create,
	},
}
