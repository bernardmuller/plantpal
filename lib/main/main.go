package main

import (
	"domain-app/internal/handlers"
	"domain-app/internal/model"
	"domain-app/internal/store/postgres"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("internal/views/*.html")),
	}
}

func mergeMaps(dst, src map[string]model.FieldError) {
	for key, value := range src {
		dst[key] = value
	}
}

func main() {
	database, err := postgres.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
	}
	plantHandler := handlers.PlantHandler{
		DB: database,
	}
	//healthcheckHandler := handlers.HealthCheckHandler{}

	e := echo.New()
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Static("/static/images", "images")
	e.Static("/static/css", "css")

	e.GET("/plants", func(c echo.Context) error {
		plants, err := handlers.PlantHandler.GetAllPlants(plantHandler, c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error fetching plants")
		}

		pageData := model.Data{Plants: plants}

		return c.Render(200, "index", model.NewPageData(pageData, model.NewFormData()))
		//return c.JSON(200, pageData)
	})

	e.POST("/plants", func(c echo.Context) error {
		fieldErrors, err := handlers.PlantHandler.CreatePlant(plantHandler, c)
		if err != nil {
			fmt.Println(fieldErrors)
			formData := model.FormData{
				Errors: map[string]string{
					"error": err.Error(),
				},
				FieldErrors: fieldErrors,
				Values: map[string]string{
					"common":         c.FormValue("common"),
					"family":         c.FormValue("family"),
					"latin":          c.FormValue("latin"),
					"category":       c.FormValue("category"),
					"origin":         c.FormValue("origin"),
					"climate":        c.FormValue("climate"),
					"tempmax":        c.FormValue("tempmax"),
					"tempmin":        c.FormValue("tempmin"),
					"ideallight":     c.FormValue("ideallight"),
					"toleratedlight": c.FormValue("toleratedlight"),
					"watering":       c.FormValue("watering"),
					"insects":        c.FormValue("insects"),
					"diseases":       c.FormValue("diseases"),
					"soil":           c.FormValue("soil"),
					"repotperiod":    c.FormValue("repotperiod"),
					"use":            c.FormValue("use"),
				}}
			fmt.Println(formData)
			return c.Render(422, "createPlantForm", formData)
		}

		return c.Redirect(302, "/plants")
	})

	e.DELETE("/plants/:id", func(c echo.Context) error {
		err := handlers.PlantHandler.DeletePlant(plantHandler, c)
		if err != nil {
			return c.String(400, err.Error())
		}

		return c.NoContent(200)
	})

	e.GET("/plants/new", func(c echo.Context) error {
		formData := model.NewFormData()
		return c.Render(200, "createPlant", formData)
	})

	//e.GET("/", func(c echo.Context) error {
	//	plants, err := handlers.PlantHandler.GetAllPlants(plantHandler, c)
	//	if err != nil {
	//		// TODO:Fix this
	//		return c.String(http.StatusNotFound, "Not Found")
	//	}
	//	fmt.Println(plants)
	//	return nil
	//	//return c.Render(200, "index", NewPageData(plants, NewFormData()))
	//})

	// router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	templates.Hello("Bernard").Render(r.Context(), w)
	// })

	//e.GET("/health-check", func(context echo.Context) error {
	//	data, err := handlers.HealthCheckHandler.ServeHTTP(healthcheckHandler, context)
	//	if err != nil {
	//		return context.JSON(http.StatusInternalServerError, map[string]string{
	//			"error": "Internal Server Error",
	//		})
	//	}
	//	fmt.Printf("%d", data)
	//	return context.JSON(http.StatusOK, data)
	//})

	port := ":8080"
	e.Logger.Fatal(e.Start(port))
}
