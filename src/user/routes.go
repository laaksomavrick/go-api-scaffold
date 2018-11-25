package user

import "github.com/laaksomavrick/goals-api/src/route"

var Routes = route.Routes{
	route.Route{
		"Index",
		"GET",
		"/users",
		Index,
	},
}
