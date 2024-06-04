package handler

import (
	"context"
	"errors"
	"github.com/bernardmuller/plantpal/services/plants/internal/service"
	plants "github.com/bernardmuller/plantpal/services/plants/plantspb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

type PlantsHandler struct {
	plantsService *service.PlantsService
	plants.UnimplementedPlantsServiceServer
}

func NewGRPCPlantsHandler(grpc *grpc.Server, plantsService *service.PlantsService) {
	grpcHandler := &PlantsHandler{
		plantsService: plantsService,
	}

	plants.RegisterPlantsServiceServer(grpc, grpcHandler)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func (h *PlantsHandler) CreatePlant(ctx context.Context, req *plants.CreatePlantRequest) (*plants.CreatePlantResponse, error) {
	plant := &plants.Plant{
		ID:             uuid.New().String(),
		Common:         req.Common,
		Family:         req.Family,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      time.Now().String(),
		Latin:          randomString(10),
		Category:       randomString(10),
		Origin:         randomString(10),
		Climate:        randomString(10),
		TempMax:        randomString(10),
		TempMin:        randomString(10),
		IdealLight:     randomString(10),
		ToleratedLight: randomString(10),
		Watering:       randomString(10),
		Insects:        randomString(10),
		Diseases:       randomString(10),
		Soil:           randomString(10),
		RepotPeriod:    randomString(10),
		Use:            randomString(10),
	}

	if len(plant.Common) == 0 {
		return nil, errors.New("Common name is required")
	}

	_, err := h.plantsService.CreatePlant(ctx, plant)
	if err != nil {
		return nil, err
	}

	res := &plants.CreatePlantResponse{
		Status: "success",
		Plant:  plant,
	}

	return res, nil
}

func (h *PlantsHandler) GetAllPlants(ctx context.Context, req *plants.GetPlantsRequest) (*plants.GetPlantsResponse, error) {
	ps, err := h.plantsService.GetAllPlants(ctx)
	if err != nil {
		return nil, err
	}

	plantsSlice := make([]*plants.Plant, len(ps))
	for i, p := range ps {
		plantsSlice[i] = &plants.Plant{
			ID:             p.ID.String(),
			Common:         p.Common,
			Family:         p.Family,
			CreatedAt:      p.CreatedAt.String(),
			UpdatedAt:      p.UpdatedAt.String(),
			Latin:          p.Latin.String,
			Category:       p.Category.String,
			Origin:         p.Origin.String,
			Climate:        p.Climate.String,
			TempMax:        p.Tempmax.String,
			TempMin:        p.Tempmin.String,
			IdealLight:     p.Ideallight.String,
			ToleratedLight: p.Toleratedlight.String,
			Watering:       p.Watering.String,
			Insects:        p.Insects.String,
			Diseases:       p.Diseases.String,
			Soil:           p.Soil.String,
			RepotPeriod:    p.Repotperiod.String,
			Use:            p.Use.String,
		}
	}

	res := &plants.GetPlantsResponse{
		Plants: plantsSlice,
	}

	return res, nil
}
