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
			Path:         "/auth/logout/google",
			Method:       "POST",
			Controller:   e.Controllers.Auth.Logout,
			RequiresAuth: true,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/auth/google",
			Method:       "GET",
			Controller:   e.Controllers.Auth.GetProvider,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
		{
			Path:         "/auth/logout",
			Method:       "POST",
			Controller:   e.Controllers.Auth.Logout,
			RequiresAuth: false,
			Validation: endpoints.Validation{
				Enable: false,
				Entity: nil,
			},
		},
	}
}

//p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
//	// try to get the user without re-authenticating
//	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
//		t, _ := template.New("foo").Parse(userTemplate)
//		t.Execute(res, gothUser)
//	} else {
//		gothic.BeginAuthHandler(res, req)
//	}
//})
