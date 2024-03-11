package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/bernardmuller/domain-app/cms_db"
	"github.com/bernardmuller/domain-app/db"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

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
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	router.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		// tmplStr := "<div>Hello There!</div>"
		w.Header().Set("Content-Type", "text/html")
		hello("bernardmuller").Render(r.Context(), w)
	})

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		db_search_param := params.Get("db")

		switch db_search_param {
		case "test":
			_, err := db.Connect_db()
			if err != nil {
				log.Println(err)
				response := Response{Ok: false, Message: "Error connecting to database."}

				js, _ := json.Marshal(response)
				w.Header().Add("Content-Type", "application/json")
				w.Write(js)
				return
			}
			response := Response{Ok: true, Message: "Database is alive!"}
			js, _ := json.Marshal(response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(js)
		case "cms":
			db, err := cms_db.Connect_cms_db()
			if err != nil {
				log.Println(err)
				response := Response{Ok: false, Message: "Error connecting to CMS database."}

				js, _ := json.Marshal(response)
				w.Header().Add("Content-Type", "application/json")
				w.Write(js)
				return
			}

			cms_db.Disconnect_cms_db(db)

			response := Response{Ok: true, Message: "CMS Database is alive!"}
			js, _ := json.Marshal(response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(js)
		default:
			response := Response{Ok: true, Message: "All good here!"}
			js, _ := json.Marshal(response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(js)
		}
	})

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
