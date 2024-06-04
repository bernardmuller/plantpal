package service

import (
	"context"
	plants "github.com/bernardmuller/plantpal/services/plants/plantspb"
	"github.com/bernardmuller/plantpal/store/postgres"
	"github.com/google/uuid"
)

type PlantsService struct {
	DB *postgres.Queries
}

func NewPlantsService(db *postgres.Queries) *PlantsService {
	return &PlantsService{
		DB: db,
	}
}

func (s *PlantsService) CreatePlant(c context.Context, plant *plants.Plant) (postgres.Plant, error) {
	parsedID, _ := uuid.Parse(plant.ID)
	params := postgres.CreatePlantParams{
		ID:     parsedID,
		Common: plant.Common,
		Family: plant.Family,
	}
	newPlant, createErr := s.DB.CreatePlant(c, params)
	if createErr != nil {
		return postgres.Plant{}, createErr
	}
	return newPlant, nil
}

func (s *PlantsService) GetAllPlants(ctx context.Context) ([]postgres.Plant, error) {
	plants, err := s.DB.GetAllPlants(ctx)
	if err != nil {
		return nil, err
	}
	return plants, nil
}
