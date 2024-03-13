package main

import (
	"domain-app/internal/handlers"
	"domain-app/internal/templates"
	"domain-app/internal/websockets"
	"flag"
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

// flag.Parse()
// 	hub := newHub()
// 	go hub.run()
// 	http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 	})

// 	server := &http.Server{
// 		Addr:              *addr,
// 		ReadHeaderTimeout: 3 * time.Second,
// 	}
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}

var upgrader = websocket.Upgrader{} // use default options

func main() {
	flag.Parse()
	hub := websockets.NewHub()
	go hub.Run()
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

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websockets.ServeWs(hub, w, r)
	})

	port := "8080"
	logger.Info("Server started", slog.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
