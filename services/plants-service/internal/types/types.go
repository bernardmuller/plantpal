package types

import (
	"context"

	plants "github.com/bernardmuller/plantpal/services/plants-service/plantspb"
)

type PlantsRepository interface {
	CreatePlant(context.Context, *plants.Plant) error
	GetAllPlants(context.Context) []*plants.Plant
}
