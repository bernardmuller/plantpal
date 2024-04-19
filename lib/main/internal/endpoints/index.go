package endpoints

import (
	"domain-app/internal/config"
	"domain-app/internal/middleware"
	"domain-app/internal/store/postgres"
	"errors"
	"github.com/labstack/echo/v4"
)

type Validation struct {
	Enable bool
	Entity interface{}
}

type Endpoint struct {
	Controller   func(c echo.Context) error
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

type Temp struct {
	UserId string
}

func protectedRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//if !ok {
		//	return c.String(http.StatusUnauthorized, `{"access": "unauthorized"}`)
		//}
		//
		//authContext := utils.CustomContext{
		//	Context: c,
		//	Data:    c.Get("Data"),
		//	CurrentUser: Temp{
		//		UserId: claims.Subject,
		//	},
		//}
		//
		//cc := authContext
		return next(c)
	}
}

func (f EndpointFactory) createEndpoint(endpoint Endpoint) error {
	if endpoint.Method != "GET" &&
		endpoint.Method != "POST" &&
		endpoint.Method != "PUT" &&
		endpoint.Method != "PATCH" &&
		endpoint.Method != "DELETE" {
		return errors.New("Invalid method on endpoint creation")
	}
	var middlewareFunctions []echo.MiddlewareFunc
	middlewareFunctions = append(middlewareFunctions, middleware.CreateCustomContext)
	if endpoint.RequiresAuth {
		middlewareFunctions = append(middlewareFunctions, protectedRoute)
	}
	//if endpoint.Validation.Enable {
	//	middlewareFunctions = append(middlewareFunctions, func(next echo.HandlerFunc) echo.HandlerFunc {
	//		return ParseEndpointParams(next, &endpoint.Validation.Entity)
	//	})
	//}
	f.ApiConfig.Router.Add(endpoint.Method, endpoint.Path, endpoint.Controller, middlewareFunctions...)
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
