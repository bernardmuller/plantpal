package controllers

import (
	"domain-app/internal/config"
	"domain-app/internal/model"
	"domain-app/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PlantsController struct {
	config *config.ApiConfig
}

func (controller *PlantsController) GetAllPlants(c echo.Context) error {
	plants, err := services.PlantsDbService{DB: controller.config.Database}.GetAllPlants(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching plants")
	}
	pageData := model.Data{Plants: plants}
	return c.Render(200, "index", model.NewPageData(pageData, model.NewFormData()))
}
