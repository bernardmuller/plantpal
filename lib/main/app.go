package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"./db"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func fly_away() string {
	return "Fly!"
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	db, err := connect_db()

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
