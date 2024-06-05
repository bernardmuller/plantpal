package handler

import (
	"encoding/json"
	"errors"
	"github.com/bernardmuller/plantpal/internal/utils"
	"github.com/bernardmuller/plantpal/services/plants/internal/service"
	plants "github.com/bernardmuller/plantpal/services/plants/plantspb"
	"github.com/google/uuid"
	"net/http"
	"time"
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

func (h *PlantsHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /plants", h.CreatePlant)
	router.HandleFunc("GET /plants", h.GetPlants)
}

func (h *PlantsHttpHandler) CreatePlant(w http.ResponseWriter, r *http.Request) {
	var plant plants.Plant

	err := json.NewDecoder(r.Body).Decode(&plant)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if len(plant.Common) == 0 {
		utils.WriteError(w, http.StatusBadRequest, errors.New("Common name is required"))
		return
	}

	if len(plant.Family) == 0 {
		utils.WriteError(w, http.StatusBadRequest, errors.New("Family name is required"))
		return
	}

	plant.ID = uuid.New().String()
	plant.CreatedAt = time.Now().String()
	plant.UpdatedAt = time.Now().String()

	_, err = h.plantsService.CreatePlant(r.Context(), &plant)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &plants.CreatePlantResponse{
		Status: "success",
		Plant:  &plant,
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *PlantsHttpHandler) GetPlants(w http.ResponseWriter, r *http.Request) {
	ps, err := h.plantsService.GetAllPlants(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
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
	utils.WriteJSON(w, http.StatusOK, res)
}
