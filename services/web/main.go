package main

import (
	"github.com/bernardmuller/plantpal/services/web/server"
)

func main() {
	httpServer := server.NewHttpServer(":8000")
	httpServer.Run()
}
