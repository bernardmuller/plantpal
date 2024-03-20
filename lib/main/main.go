package main

import (
	"domain-app/internal/handlers"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type Plant struct {
	Name string
}

func main() {
	plantHandler := handlers.PlantHandler{}
	healthcheckHandler := handlers.HealthCheckHandler{}

	e := echo.New()
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/plants", func(c echo.Context) error {
		err := handlers.PlantHandler.GetAllPlants(plantHandler, c)
		if err != nil {
			fmt.Println("Error:", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.NoContent(http.StatusOK)
	})

	e.GET("/", func(c echo.Context) error {
		plant := Plant{Name: "Fiddle leaf fig"}
		return c.Render(200, "index", plant)
	})

	// router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	templates.Hello("Bernard").Render(r.Context(), w)
	// })

	e.GET("/health-check", func(context echo.Context) error {
		data, err := handlers.HealthCheckHandler.ServeHTTP(healthcheckHandler, context)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Internal Server Error",
			})
		}
		fmt.Printf("%d", data)
		return context.JSON(http.StatusOK, data)
	})

	port := ":8080"
	e.Logger.Fatal(e.Start(port))
}
