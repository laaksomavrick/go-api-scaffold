package user

import "github.com/laaksomavrick/goals-api/src/core"

var Routes = core.Routes{
	core.Route{
		"Index",
		"GET",
		"/users",
		Index,
	},
}
