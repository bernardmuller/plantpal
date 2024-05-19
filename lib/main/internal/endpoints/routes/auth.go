package routes

import (
	"domain-app/internal/endpoints"
)

func (e *Endpoints) AuthEndpoints() []endpoints.Endpoint {
	return []endpoints.Endpoint{
		{
			Path:         "/auth/google/callback",
			Method:       "GET",
			Controller:   e.Controllers.Auth.GetCallback,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/auth/logout",
			Method:       "GET",
			Controller:   e.Controllers.Auth.Logout,
			RequiresAuth: true,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/auth/login",
			Method:       "GET",
			Controller:   e.Controllers.Auth.GetProvider,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
	}
}
