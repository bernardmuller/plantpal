package routes

import (
	"domain-app/internal/controllers"
	"domain-app/internal/endpoints"
)

type Endpoints struct {
	Controllers controllers.Controllers
	//validators interface{}
}

func (e *Endpoints) Endpoints() []endpoints.Endpoint {
	return []endpoints.Endpoint{
		{
			Path:               "/plants",
			Method:             "GET",
			Controller:         e.Controllers.Plants.GetAllPlants,
			RequiresAuth:       false,
			RequiresValidation: false,
		},
		//{
		//	Path:   "/plants/new",
		//	Method: "GET",
		//	Controller: func(c echo.Context) error {
		//		formData := model.NewFormData()
		//		return c.Render(200, "createPlant", formData)
		//	},
		//	RequiresAuth:       false,
		//	RequiresValidation: false,
		//},
		//{
		//	Path:   "/plants",
		//	Method: "POST",
		//	Controller: func(c echo.Context) error {
		//		fieldErrors, err := services.PlantController.CreatePlant(plantController, c)
		//		if err != nil {
		//			formData := model.FormData{
		//				Errors: map[string]string{
		//					"error": err.Error(),
		//				},
		//				FieldErrors: fieldErrors,
		//				Values: map[string]string{
		//					"common":         c.FormValue("common"),
		//					"family":         c.FormValue("family"),
		//					"latin":          c.FormValue("latin"),
		//					"category":       c.FormValue("category"),
		//					"origin":         c.FormValue("origin"),
		//					"climate":        c.FormValue("climate"),
		//					"tempmax":        c.FormValue("tempmax"),
		//					"tempmin":        c.FormValue("tempmin"),
		//					"ideallight":     c.FormValue("ideallight"),
		//					"toleratedlight": c.FormValue("toleratedlight"),
		//					"watering":       c.FormValue("watering"),
		//					"insects":        c.FormValue("insects"),
		//					"diseases":       c.FormValue("diseases"),
		//					"soil":           c.FormValue("soil"),
		//					"repotperiod":    c.FormValue("repotperiod"),
		//					"use":            c.FormValue("use"),
		//				}}
		//			fmt.Println(formData)
		//			return c.Render(422, "createPlantForm", formData)
		//		}
		//
		//		return c.Redirect(302, "/plants")
		//	},
		//	RequiresAuth:       false,
		//	RequiresValidation: false,
		//},
		//{
		//	Path:   "/plants/:id",
		//	Method: "DELETE",
		//	Controller: func(c echo.Context) error {
		//		err := services.PlantController.DeletePlant(plantController, c)
		//		if err != nil {
		//			return c.String(400, err.Error())
		//		}
		//
		//		return c.NoContent(200)
		//	},
		//	RequiresAuth:       false,
		//	RequiresValidation: false,
		//},
	}
}
