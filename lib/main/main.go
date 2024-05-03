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
	factory := &endpoints.EndpointFactory{
		ApiConfig: apiConfig,
		Endpoints: append(apiEndpoints.PlantEndpoints(), apiEndpoints.AuthEndpoints()...),
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

	//handler := cors.Default().Handler(mux)
	//http.ListenAndServe(":8080", handler)
	//server := apiConfig.Router.Start(apiConfig.PORT)
	apiConfig.Router.Logger.Fatal(apiConfig.Router.Start(apiConfig.PORT))

}
