package main

import gopher/core

func main() {
	e, services, err := core.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	core.InitAll(services)
	services.Jobs.Start()
	core.Route.Register("GET", "/", homeHandler, "home", "core")

	e.Logger.Fatal(e.Start(":8080"))
}
