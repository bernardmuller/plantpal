package endpoints

import (
	"domain-app/internal/config"
	"domain-app/internal/middleware"
	"domain-app/internal/store/postgres"
	"errors"
	"net/http"
)

type Validation struct {
	Enable bool
	Entity interface{}
}

type Endpoint struct {
	Controller   func(w http.ResponseWriter, r *http.Request)
	Method       string
	Path         string
	Validation   Validation
	RequiresAuth bool
}

type Data struct {
	Plants []postgres.Plant
}

func NewPageData(data Data, form FormData) PageData {
	return PageData{
		Data: data,
		Form: form,
	}
}

type FormData struct {
	Errors map[string]string
	Values map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

type PageData struct {
	Data Data
	Form FormData
}

// Busy working on a way to implement dependency injection on the handlers
// as well as moving the endpoints to a factory pattern
// so that the handlers can be injected into the endpoints along with the echo.Context and DB, etc.
// We will be able to set the endpoints with spesifications on which requires auth and which does not
//
// apiConfig => endpoints factory => endpoints => json handlers
//
//	|															=> html handlers
//	|> validate requests
//	|> validate auth

type EndpointFactory struct {
	ApiConfig *config.ApiConfig
	Endpoints []Endpoint
}

//func ParseEndpointParams(next echo.HandlerFunc, entity *interface{}) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		customContext := c.(utils.CustomContext)
//		parsed := utils.PopulateStructFromForm(customContext.Request(), entity)
//		customContext.SetData(parsed)
//
//		return next(customContext)
//	}
//}

//func Authorise() {
//	return clerkhttp.WithHeaderAuthorization()
//}

func (f EndpointFactory) createEndpoint(endpoint Endpoint) error {

	switch endpoint.Method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
	default:
		return errors.New("Invalid method on endpoint creation")
	}

	var middlewareFunctions []func(http.Handler) http.Handler
	middlewareFunctions = append(middlewareFunctions, middleware.CreateCustomContext)

	if endpoint.RequiresAuth {
		// Add authentication middleware
	}

	// Add validation middleware if enabled
	if endpoint.Validation.Enable {
		// Add validation middleware
	}

	// Chain the middleware functions
	var handler http.Handler = http.HandlerFunc(endpoint.Controller)
	for _, middlewareFunc := range middlewareFunctions {
		handler = middlewareFunc(handler)
	}

	// Add the endpoint to the router
	f.ApiConfig.Router.HandleFunc(endpoint.Path, endpoint.Controller)
	return nil
}

func (f EndpointFactory) CreateEndpoints() {
	for _, endpoint := range f.Endpoints {
		err := f.createEndpoint(endpoint)
		if err != nil {
			panic(err)
		}
	}
}
