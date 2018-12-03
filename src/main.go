package main

import (
	"github.com/laaksomavrick/goals-api/src/config"
	"github.com/laaksomavrick/goals-api/src/core"
	"github.com/laaksomavrick/goals-api/src/healthz"
	"github.com/laaksomavrick/goals-api/src/user"
)

func main() {

	// load all the required env values
	config := config.NewConfig()

	// initialize the server object
	// values in this struct are available to all handlers
	server := core.NewServer(core.NewRouter(), core.NewDatabase(config), config)
	// initialize exported routes from packages
	routes := append(
		healthz.Routes,
		user.Routes...,
	)
	// initialize the application given our routes
	server.Init(routes)

}
