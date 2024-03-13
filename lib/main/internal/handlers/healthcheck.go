package handlers

import (
	"domain-app/internal/store/cms_db"
	"domain-app/internal/store/db"
	"encoding/json"
	"log"
	"net/http"
)

type HealthCheck struct{}

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func HealthCheckHandler() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

}
