package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	plants "github.com/bernardmuller/plantpal/services/plants-service/plantspb"
)

type httpServer struct {
	addr string
}

type Host string

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}

	return conn
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	//router := http.NewServeMux()
	router := mux.NewRouter()

	var host Host
	if os.Getenv("ENV") != "production" {
		host = "127.0.0.1" // Use localhost for production
	} else {
		host = "0.0.0.0" // Use any available address for non-production
	}

	conStr := fmt.Sprintf("%s:9001", host)
	conn := NewGRPCClient(conStr)
	defer conn.Close()
	//m := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := plants.NewPlantsServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		ps, err := c.GetAllPlants(ctx, &plants.GetPlantsRequest{})
		if err != nil {
			fmt.Println("client error: %v", err)
		}

		t := template.Must(template.ParseFiles("services/web/views/layouts/main.html", "services/web/views/pages/plants.tmpl.html"))

		if err := t.Execute(w, ps.Plants); err != nil {
			fmt.Println("template error: %v", err)
		}
	})

	router.HandleFunc("/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
		c := plants.NewPlantsServiceClient(conn)

		vars := mux.Vars(r)
		plantId := vars["id"]

		fmt.Println("plantId: ", plantId)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		ps, err := c.GetAllPlants(ctx, &plants.GetPlantsRequest{})
		if err != nil {
			fmt.Println("client error: %v", err)
		}

		var plant *plants.Plant
		for _, p := range ps.Plants {
			if p.ID == plantId {
				plant = p
				break
			}
		}
		t := template.Must(template.ParseFiles("services/web/views/layouts/main.html", "services/web/views/pages/plant.tmpl.html"))

		if err := t.Execute(w, plant); err != nil {
			fmt.Println("template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
