package controllers

import (
	"domain-app/internal/config"
	"domain-app/internal/services"
)

type PlantsController struct {
	config      *config.ApiConfig
	authService *services.AuthDBService
}

type AuthController struct {
	config      *config.ApiConfig
	userService *services.UserDBService
	authService *services.AuthDBService
}

type UsersController struct {
	config      *config.ApiConfig
	userService *services.UserDBService
	authService *services.AuthDBService
}

type Controllers struct {
	Plants *PlantsController
	Auth   *AuthController
	Users  *UsersController
}

func NewControllers(ac config.ApiConfig) Controllers {
	userService := services.UserDBService{DB: ac.Database}
	authService := services.AuthDBService{DB: ac.Database}

	return Controllers{
		Plants: &PlantsController{
			config: &ac,
		},
		Auth: &AuthController{
			config:      &ac,
			userService: &userService,
			authService: &authService,
		},
	}
}
