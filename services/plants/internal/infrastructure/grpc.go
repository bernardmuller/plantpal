package infrastructure

import (
	"github.com/bernardmuller/plantpal/internal/module"
	"github.com/bernardmuller/plantpal/services/plants/internal/handler"
	"github.com/bernardmuller/plantpal/services/plants/internal/service"
	"github.com/bernardmuller/plantpal/store/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	addr string
	DB   *postgres.Queries
}

func NewGrpcServer(config *module.ModuleConfig) *Server {
	return &Server{addr: config.PORT.GRPC, DB: config.Database}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	plantsService := service.NewPlantsService(s.DB)
	handler.NewGRPCPlantsHandler(grpcServer, plantsService)

	log.Println("Server is running on port", s.addr)

	return grpcServer.Serve(lis)
}
