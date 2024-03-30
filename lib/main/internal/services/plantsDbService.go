package services

import (
	"context"
	"domain-app/internal/model"
	"domain-app/internal/store/postgres"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IPlantService interface {
	GetAllPlants(c echo.Context) ([]postgres.Plant, error)
	GetPlantByID(c echo.Context) (postgres.Plant, error)
	GetPlantByCommon(c echo.Context) (postgres.Plant, error)
	CreatePlant(c echo.Context) (map[string]model.FieldError, error)
	DeletePlant(c echo.Context) error
}

type PlantsDbService struct {
	DB *postgres.Queries
}

func (service PlantsDbService) GetAllPlants(c context.Context) ([]postgres.Plant, error) {
	plants, err := service.DB.GetAllPlants(c)
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service PlantsDbService) GetPlantByID(c context.Context, id string) (postgres.Plant, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return postgres.Plant{}, err
	}

	plant, plantErr := service.DB.GetPlantByID(c, uuid)
	if plantErr != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (service PlantsDbService) GetPlantByCommon(c context.Context, common string) (postgres.Plant, error) {
	if common == "" {
		return postgres.Plant{}, errors.New("missing common name")
	}

	plant, err := service.DB.GetPlantByCommon(c, common)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (service PlantsDbService) CreatePlant(c context.Context) (map[string]model.FieldError, error) {
	//params := postgres.CreatePlantParams{
	//	Common:         c.FormValue("common"),
	//	Family:         c.FormValue("family"),
	//	Latin:          sql.NullString{String: c.FormValue("latin"), Valid: true},
	//	Category:       sql.NullString{String: c.FormValue("category"), Valid: true},
	//	Origin:         sql.NullString{String: c.FormValue("origin"), Valid: true},
	//	Climate:        sql.NullString{String: c.FormValue("climate"), Valid: true},
	//	Tempmax:        sql.NullString{String: c.FormValue("tempmax"), Valid: true},
	//	Tempmin:        sql.NullString{String: c.FormValue("tempmin"), Valid: true},
	//	Ideallight:     sql.NullString{String: c.FormValue("ideallight"), Valid: true},
	//	Toleratedlight: sql.NullString{String: c.FormValue("toleratedlight"), Valid: true},
	//	Watering:       sql.NullString{String: c.FormValue("watering"), Valid: true},
	//	Insects:        sql.NullString{String: c.FormValue("insects"), Valid: true},
	//	Diseases:       sql.NullString{String: c.FormValue("diseases"), Valid: true},
	//	Soil:           sql.NullString{String: c.FormValue("soil"), Valid: true},
	//	Repotperiod:    sql.NullString{String: c.FormValue("repotperiod"), Valid: true},
	//	Use:            sql.NullString{String: c.FormValue("use"), Valid: true},
	//}
	//
	//paramsSlice := reflect.ValueOf(params)
	//fieldErrors := make([]model.FieldError, paramsSlice.NumField())
	//
	//for i := 0; i < paramsSlice.NumField(); i++ {
	//	n := paramsSlice.Type().Field(i).Name
	//	if paramsSlice.Field(i).Interface() == "" {
	//		newFieldError := model.FieldError{Field: n, Message: "Please enter a valid " + n + " name"}
	//		fieldErrors = append(fieldErrors, newFieldError)
	//	}
	//}
	//
	//fieldErrorsMap := make(map[string]model.FieldError)
	//for _, err := range fieldErrors {
	//	fieldErrorsMap[err.Field] = err
	//}
	//
	//if len(fieldErrorsMap) > 1 {
	//	return fieldErrorsMap, errors.New("Please enter valid values for all fields")
	//}
	//
	//dbPlant, err := service.DB.GetPlantByCommon(c, params.Common)
	//fmt.Println(dbPlant)
	//if err == nil {
	//	fmt.Println("Plant with that name already exists.")
	//	fieldErrorsMap["Common"] = model.FieldError{Field: "Common", Message: "Plant with that name already exists."}
	//	return fieldErrorsMap, errors.New("Plant with that name already exists.")
	//}
	//
	//_, createErr := service.DB.CreatePlant(c, params)
	//if createErr != nil {
	//	return nil, createErr
	//}
	return nil, nil
}

func (service PlantsDbService) DeletePlant(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	_, err = service.DB.DeletePlant(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}
