package infrastructure

import (
	"github.com/bernardmuller/plantpal/internal/module"
	"github.com/bernardmuller/plantpal/services/plants/internal/handler"
	"github.com/bernardmuller/plantpal/services/plants/internal/service"
	"github.com/bernardmuller/plantpal/store/postgres"
	"log"
	"net/http"
)

type httpServer struct {
	addr string
	DB   *postgres.Queries
}

func NewHttpServer(config *module.ModuleConfig) *httpServer {
	return &httpServer{addr: config.PORT.HTTP, DB: config.Database}
}

func (s *httpServer) Start() error {
	router := http.NewServeMux()

	plantService := service.NewPlantsService(s.DB)
	plantHandler := handler.NewHttpPlantsHandler(plantService)
	plantHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
