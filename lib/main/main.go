package main

import (
	"domain-app/internal/handlers"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"os"

	"github.com/labstack/echo"
)

func main() {
	flag.Parse()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	fmt.Println(os.Getenv("TEST_DATABASE_URL"))

	plantHandler := handlers.PlantHandler{}
	healthcheckHandler := handlers.HealthCheckHandler{}

	e := echo.New()

	e.GET("/plants", func(c echo.Context) error {
		err := handlers.PlantHandler.GetAllPlants(plantHandler, c)
		if err != nil {
			fmt.Println("Error")
		}
		return nil
	})

	//router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	content := templates.GuestIndex()
	//	templates.Layout(content).Render(r.Context(), w)
	//})

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

	port := "8080"
	logger.Info("Server started", slog.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, e))
}
