package auth

import "github.com/laaksomavrick/goals-api/src/core"

var Routes = core.Routes{
	core.Route{
		Name:        "Authenticate",
		Method:      "POST",
		Pattern:     "/auth",
		HandlerFunc: Create,
	},
}
