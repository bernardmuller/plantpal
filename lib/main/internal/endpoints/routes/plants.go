package routes

import (
	"domain-app/internal/controllers"
	"domain-app/internal/endpoints"
	"domain-app/internal/store/postgres"
)

type Endpoints struct {
	Controllers controllers.Controllers
}

func (e *Endpoints) PlantEndpoints() []endpoints.Endpoint {
	var createPlantParams postgres.CreatePlantParams
	return []endpoints.Endpoint{
		{
			Path:         "/plants",
			Method:       "GET",
			Controller:   e.Controllers.Plants.GetAllPlants,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/plants/new",
			Method:       "GET",
			Controller:   e.Controllers.Plants.GetCreatePlantForm,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/plants",
			Method:       "POST",
			Controller:   e.Controllers.Plants.CreatePlant,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: true,
				Entity: &createPlantParams,
			},
		},
		{
			Path:         "/plants/:id",
			Method:       "DELETE",
			Controller:   e.Controllers.Plants.DeletePlant,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
	}
}
