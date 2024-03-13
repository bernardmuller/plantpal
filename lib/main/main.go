package main

import (
	"domain-app/internal/handlers"
	"domain-app/internal/templates"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

var upgrader = websocket.Upgrader{} // use default options

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

	router.HandleFunc("/ws", echo)

	port := "8080"
	logger.Info("Server started", slog.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
