package endpoints

import (
	"domain-app/internal/store/postgres"
	"github.com/labstack/echo/v4"
)

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

var PlantEndpoints = []func(c *echo.Context) error{
	//handlers.PlantHandler{}GetAllPlants,
}

func CreatePlantsEndpoints(apiConfig interface{}) []func(c *echo.Context) error {
	//return
}
