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

	apiConfig.Router.Logger.Fatal(apiConfig.Router.Start(apiConfig.PORT))
	//
	//p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
	//
	//	user, err := gothic.CompleteUserAuth(res, req)
	//	if err != nil {
	//		fmt.Fprintln(res, err)
	//		return
	//	}
	//	t, _ := template.New("foo").Parse(userTemplate)
	//	t.Execute(res, user)
	//})
	//
	//p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
	//	gothic.Logout(res, req)
	//	res.Header().Set("Location", "/")
	//	res.WriteHeader(http.StatusTemporaryRedirect)
	//})
	//
	//p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
	//	// try to get the user without re-authenticating
	//	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
	//		t, _ := template.New("foo").Parse(userTemplate)
	//		t.Execute(res, gothUser)
	//	} else {
	//		gothic.BeginAuthHandler(res, req)
	//	}
	//})

}
