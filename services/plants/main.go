package main

import (
	"github.com/bernardmuller/plantpal/services/plants/internal/config"
	"github.com/bernardmuller/plantpal/services/plants/internal/infrastructure"
)

func main() {
	moduleConfig := config.CreateConfig()

	httpServer := infrastructure.NewHttpServer(moduleConfig)
	go httpServer.Start()

	grpcServer := infrastructure.NewGrpcServer(":9000")
	grpcServer.Start()
}
