package endpoints

import "github.com/labstack/echo/v4"

type Endpoints struct {
	Plants []func(c *echo.Context) error
}

func CreateEndpoints() *Endpoints {
	return &Endpoints{
		Plants: PlantEndpoints,
	}
}
