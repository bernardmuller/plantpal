package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type PlantHandler struct{}

func (handler PlantHandler) GetAllPlants(c echo.Context) error {
	fmt.Println("All plants")
	return nil
}
