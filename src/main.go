package main

import (
	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/healthz"
	"github.com/laaksomavrick/goals-api/src/user"
)

func main() {

	// initialize the server object
	// values in this struct are available to all handlers
	server := core.NewServer(core.NewRouter(), core.NewDatabase())
	// initialize exported routes from packages
	routes := append(
		healthz.Routes,
		user.Routes...,
	)
	// initialize the application given our routes
	server.Init(routes)

}
