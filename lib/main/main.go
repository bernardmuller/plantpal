package main

import (
	"domain-app/internal/handlers"
	"domain-app/internal/templates"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func fly_away() string {
	return "Fly!"
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := templates.GuestIndex()
		templates.Layout(content).Render(r.Context(), w)
	})

	router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		templates.Hello("Bernard").Render(r.Context(), w)
	})

	router.HandleFunc("/health-check", handlers.HealthCheckHandler().ServeHTTP)

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
