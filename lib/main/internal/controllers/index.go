package controllers

import "domain-app/internal/config"

type Controllers struct {
	Plants *PlantsController
}

func NewControllers(ac config.ApiConfig) Controllers {
	return Controllers{
		Plants: &PlantsController{
			config: &ac,
		},
	}
}
