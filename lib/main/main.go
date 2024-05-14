package main

import (
	"domain-app/internal/auth"
	"domain-app/internal/config"
	"domain-app/internal/controllers"
	"domain-app/internal/endpoints"
	"domain-app/internal/endpoints/routes"
)

func main() {

	auth.NewAuth()

	apiConfig := config.CreateAPIConfig()
	apiControllers := controllers.NewControllers(*apiConfig)
	apiEndpoints := routes.Endpoints{
		Controllers: apiControllers,
	}

	endpointFuncs := []func() []endpoints.Endpoint{
		apiEndpoints.AuthEndpoints,
		apiEndpoints.PlantEndpoints,
		apiEndpoints.UsersEndpoints,
	}

	factory := &endpoints.EndpointFactory{
		ApiConfig: apiConfig,
		Endpoints: []endpoints.Endpoint{},
	}

	for _, endpointFunc := range endpointFuncs {
		factory.Endpoints = append(factory.Endpoints, endpointFunc()...)
	}

	factory.CreateEndpoints()

	apiConfig.Router.Logger.Fatal(apiConfig.Router.Start(apiConfig.PORT))
}
