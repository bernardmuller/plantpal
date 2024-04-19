package routes

import (
	"domain-app/internal/endpoints"
)

func (e *Endpoints) AuthEndpoints() []endpoints.Endpoint {
	return []endpoints.Endpoint{
		{
			Path:         "/auth/login",
			Method:       "GET",
			Controller:   e.Controllers.Auth.GetLoginPage,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
	}
}
