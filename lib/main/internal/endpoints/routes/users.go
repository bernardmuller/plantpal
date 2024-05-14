package routes

import (
	"domain-app/internal/endpoints"
)

func (e *Endpoints) UsersEndpoints() []endpoints.Endpoint {
	return []endpoints.Endpoint{
		{
			Path:         "/users/{id}",
			Method:       "GET",
			Controller:   e.Controllers.Users.GetUserById,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/users/session",
			Method:       "GET",
			Controller:   e.Controllers.Users.GetUserBySessionId,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
	}
}
