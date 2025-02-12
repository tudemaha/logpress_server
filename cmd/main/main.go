package main

import (
	"github.com/tudemaha/logpress_server/pkg/server"
	"github.com/tudemaha/logpress_server/routes"
)

func main() {
	routes.LoadRoutes()
	server.StartServer()
}
