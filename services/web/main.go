package main

import (
	"fmt"
	"github.com/bernardmuller/plantpal/services/web/server"
)

func main() {
	httpServer := server.NewHttpServer(":8000")
	httpServer.Run()
	fmt.Println("Cool Server!")
}
