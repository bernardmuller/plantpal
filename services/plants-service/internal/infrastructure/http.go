package infrastructure

import (
	"log"
	"net/http"

	"github.com/bernardmuller/plantpal/internal/module"
	"github.com/bernardmuller/plantpal/services/plants-service/internal/handler"
	"github.com/bernardmuller/plantpal/services/plants-service/internal/service"
	"github.com/bernardmuller/plantpal/store/postgres"
)

type httpServer struct {
	addr string
	DB   *postgres.Queries
}

// CORS middleware
func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                                                // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                                                 // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin") // Allowed headers

		// If it's a preflight request, stop here
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		handler.ServeHTTP(w, r)
	})
}

func NewHttpServer(config *module.ModuleConfig) *httpServer {
	return &httpServer{addr: config.PORT.HTTP, DB: config.Database}
}

func (s *httpServer) Start() error {
	router := http.NewServeMux()

	plantService := service.NewPlantsService(s.DB)
	plantHandler := handler.NewHttpPlantsHandler(plantService)
	plantHandler.RegisterRouter(router)
	// Wrap the router with the CORS middleware
	corsRouter := corsMiddleware(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, corsRouter)
}
