package main

import (
	"domain-app/internal/handlers"
	"domain-app/internal/templates"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

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

	port := "8080"
	logger.Info("Server started", slog.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
