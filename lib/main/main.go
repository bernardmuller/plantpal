package main

import (
	"domain-app/internal/config"
	"domain-app/internal/controllers"
	"domain-app/internal/endpoints"
	"domain-app/internal/endpoints/routes"
)

func main() {
	apiConfig := config.CreateAPIConfig()
	apiControllers := controllers.NewControllers(*apiConfig)
	apiEndpoints := routes.Endpoints{
		Controllers: apiControllers,
	}
	factory := &endpoints.EndpointFactory{
		ApiConfig: apiConfig,
		Endpoints: apiEndpoints.Endpoints(),
	}
	factory.CreateEndpoints()

	//e.GET("/health-check", func(context echo.Context) error {
	//	data, err := Controllers.HealthCheckController.ServeHTTP(healthcheckController, context)
	//	if err != nil {
	//		return context.JSON(http.StatusInternalServerError, map[string]string{
	//			"error": "Internal Server Error",
	//		})
	//	}
	//	fmt.Printf("%d", data)
	//	return context.JSON(http.StatusOK, data)
	//})

	apiConfig.Router.Logger.Fatal(apiConfig.Router.Start(apiConfig.PORT))
}
