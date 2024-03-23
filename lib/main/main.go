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

func NewPageData(data Data, form FormData) PageData {
	return PageData{
		Data: data,
		Form: form,
	}
}

type Data struct {
	Plants []postgres.Plant
}

func NewData() model.Plants {
	return model.Plants{
		Plants: []model.Plant{
			{
				Common: "Pothos",
				Family: "Araceae",
				Id:     1,
			},
			{
				Common: "Cactus",
				Family: "Cactaceae",
				Id:     2,
			},
		},
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
			fmt.Println("Error:", err)
			return c.String(http.StatusInternalServerError, "Error fetching plants")
		}
		pageData := Data{Plants: plants}

		fmt.Sprintf("%v", pageData)
		//return c.Render(200, "index", NewPageData(pageData, NewFormData()))
		return c.JSON(200, pageData)
	})

	//e.POST("/plants", func(c echo.Context) error {
	//
	//	newPlant := NewPlant(name, family, id)
	//	data.Plants = append(data.Plants, newPlant)
	//
	//	formData := NewFormData()
	//	_ = c.Render(200, "addPlantForm", formData)
	//	return c.Render(200, "oob-plant", newPlant)
	//})

	//e.DELETE("/plants/:id", func(c echo.Context) error {
	//	idStr := c.Param("id")
	//	id, err := strconv.Atoi(idStr)
	//
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		return c.String(400, "Id must be an integer")
	//	}
	//
	//	deleted := false
	//	for i, plant := range data.Plants {
	//		if plant.Id == id {
	//			data.Plants = append(data.Plants[:i], data.Plants[i+1:]...)
	//			deleted = true
	//			break
	//		}
	//	}
	//
	//	if !deleted {
	//		return c.String(400, "Plant not found")
	//	}
	//
	//	return c.NoContent(200)
	//})

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
