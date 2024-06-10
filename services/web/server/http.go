package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/bernardmuller/plantpal/services/plants/plantspb"
)

type httpServer struct {
	addr string
}

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	//router := http.NewServeMux()
	router := mux.NewRouter()

	conn := NewGRPCClient(":9001")
	defer conn.Close()
	//m := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := plants.NewPlantsServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		ps, err := c.GetAllPlants(ctx, &plants.GetPlantsRequest{})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.ParseFiles("services/web/views/layouts/main.html", "services/web/views/pages/plants.tmpl.html"))

		if err := t.Execute(w, ps.Plants); err != nil {
			log.Fatalf("template error: %v", err)
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
			log.Fatalf("client error: %v", err)
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
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Plants</title>
</head>
<body>
    <h1>Plants List</h1>
    <table border="1">
        <tr>
            <th>Plant ID</th>
            <th>Common Name</th>
            <th>Family</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.Common}}</td>
            <td>{{.Family}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
