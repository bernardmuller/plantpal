package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bernardmuller/plantpal/services/plants-service/internal/service"
	plants "github.com/bernardmuller/plantpal/services/plants-service/plantspb"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PlantsHttpHandler struct {
	plantsService service.PlantsService
}

func NewHttpPlantsHandler(plantService *service.PlantsService) *PlantsHttpHandler {
	var ps service.PlantsService
	ps = *plantService
	return &PlantsHttpHandler{
		plantsService: ps,
	}
}

func (h *PlantsHttpHandler) RegisterRouter(router *echo.Echo) {
	router.POST("/plants", h.CreatePlant)
	router.GET("/plants", h.GetPlants)
	router.GET("/plants/:id", h.GetPlantById)
}

func (h *PlantsHttpHandler) CreatePlant(c echo.Context) error {
	var plant plants.Plant

	err := json.NewDecoder(c.Request().Body).Decode(&plant)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error decoding request body plants")
	}

	if len(plant.Common) == 0 {
		return c.String(http.StatusBadRequest, "Common name is required")
	}

	if len(plant.Family) == 0 {
		return c.String(http.StatusBadRequest, "Family name is required")
	}

	plant.ID = uuid.New().String()
	plant.CreatedAt = time.Now().String()
	plant.UpdatedAt = time.Now().String()

	_, err = h.plantsService.CreatePlant(c.Request().Context(), &plant)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := &plants.CreatePlantResponse{
		Status: "success",
		Plant:  &plant,
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *PlantsHttpHandler) GetPlants(c echo.Context) error {
	ps, err := h.plantsService.GetAllPlants(c.Request().Context())
	if err != nil {
		return c.String(http.StatusNoContent, err.Error())
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
	// utils.WriteJSON(w, http.StatusOK, res)
	return c.JSON(http.StatusOK, res)
}

func (h *PlantsHttpHandler) GetPlantById(c echo.Context) error {
	plantIdStr := c.Param("id")

	plantId, err := uuid.Parse(plantIdStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	p, err := h.plantsService.GetPlantById(c.Request().Context(), plantId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
