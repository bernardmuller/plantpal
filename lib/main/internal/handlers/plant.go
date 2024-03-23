package handlers

import (
	"domain-app/internal/store/postgres"
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
