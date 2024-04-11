package services

import (
	"context"
	"domain-app/internal/store/postgres"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IPlantService interface {
	GetAllPlants(c echo.Context) ([]postgres.Plant, error)
	GetPlantByID(c echo.Context) (postgres.Plant, error)
	GetPlantByCommon(c echo.Context) (postgres.Plant, error)
	CreatePlant(c echo.Context) (postgres.Plant, error)
	DeletePlant(c echo.Context) error
}

type PlantsDbService struct {
	DB *postgres.Queries
}

func (service PlantsDbService) GetAllPlants(c context.Context) ([]postgres.Plant, error) {
	plants, err := service.DB.GetAllPlants(c)
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service PlantsDbService) GetPlantByID(c context.Context, id string) (postgres.Plant, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return postgres.Plant{}, err
	}

	plant, plantErr := service.DB.GetPlantByID(c, uuid)
	if plantErr != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (service PlantsDbService) GetPlantByCommon(c context.Context, common string) (postgres.Plant, error) {
	if common == "" {
		return postgres.Plant{}, errors.New("missing common name")
	}

	plant, err := service.DB.GetPlantByCommon(c, common)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (service PlantsDbService) CreatePlant(c context.Context, params postgres.CreatePlantParams) (postgres.Plant, error) {
	newPlant, createErr := service.DB.CreatePlant(c, params)
	if createErr != nil {
		return postgres.Plant{}, createErr
	}
	return newPlant, nil
}

func (service PlantsDbService) DeletePlant(c context.Context, id uuid.UUID) error {
	_, err := service.DB.DeletePlant(c, id)
	if err != nil {
		return err
	}
	return nil
}
