package main

import (
	"github.com/bernardmuller/plantpal/internal/module"
	"github.com/bernardmuller/plantpal/services/plants-service/internal/infrastructure"
)

func main() {
	port := module.PORT{
		HTTP: ":8001",
		GRPC: ":9001",
	}

	moduleConfig, err := module.CreateConfig(port)
	if err != nil {
		panic(err)
	}

	httpServer := infrastructure.NewHttpServer(moduleConfig)
	go httpServer.Start()

	grpcServer := infrastructure.NewGrpcServer(moduleConfig)
	grpcServer.Start()
}
