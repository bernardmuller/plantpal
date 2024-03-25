package handlers

import (
	"domain-app/internal/store/postgres"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PlantHandler struct {
	DB *postgres.Queries
}

func (handler PlantHandler) GetAllPlants(c echo.Context) ([]postgres.Plant, error) {
	plants, err := handler.DB.GetAllPlants(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (handler PlantHandler) GetPlantByID(c echo.Context) (postgres.Plant, error) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return postgres.Plant{}, err
	}

	plant, err := handler.DB.GetPlantByID(c.Request().Context(), uuid)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (handler PlantHandler) GetPlantByCommon(c echo.Context) (postgres.Plant, error) {
	common := c.QueryParam("common")
	if common == "" {
		return postgres.Plant{}, errors.New("missing common name")
	}

	plant, err := handler.DB.GetPlantByCommon(c.Request().Context(), common)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (handler PlantHandler) CreatePlant(c echo.Context) (postgres.Plant, error) {
	name := c.FormValue("name")
	family := c.FormValue("family")

	_, err := handler.DB.GetPlantByCommon(c.Request().Context(), name)
	if err == nil {
		return postgres.Plant{}, errors.New("plant already exists")
	}

	params := postgres.CreatePlantParams{
		Common: name,
		Family: family,
		ID:     uuid.New(),
	}

	newPlant, createErr := handler.DB.CreatePlant(c.Request().Context(), params)
	if createErr != nil {
		return postgres.Plant{}, createErr
	}
	return newPlant, nil
}

func (handler PlantHandler) DeletePlant(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	_, err = handler.DB.DeletePlant(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}
