package endpoints

import (
	"domain-app/internal/config"
	"domain-app/internal/store/postgres"
	"errors"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	Controller         func(c echo.Context) error
	Method             string
	Path               string
	RequiresValidation bool
	RequiresAuth       bool
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
// 							|															=> html handlers
// 							|> validate requests
//							|> validate auth

type EndpointFactory struct {
	ApiConfig *config.ApiConfig
	Endpoints []Endpoint
}

func (f EndpointFactory) createEndpoint(endpoint Endpoint) error {
	if endpoint.Method != "GET" && endpoint.Method != "POST" && endpoint.Method != "PUT" && endpoint.Method != "DELETE" {
		return errors.New("Invalid method on endpoint creation")
	}
	f.ApiConfig.Router.Add(endpoint.Method, endpoint.Path, func(c echo.Context) error {
		endpoint.Controller(c)
		return nil
	})
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
