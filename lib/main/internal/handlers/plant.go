package handlers

import (
	"database/sql"
	"domain-app/internal/model"
	"domain-app/internal/store/postgres"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"reflect"
)

type PlantHandler struct {
	DB *postgres.Queries
}

func (handler PlantHandler) GetAllPlants(c echo.Context) ([]postgres.Plant, error) {
	plants, err := handler.DB.GetAllPlants(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (handler PlantHandler) GetPlantByID(c echo.Context) (postgres.Plant, error) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return postgres.Plant{}, err
	}

	plant, err := handler.DB.GetPlantByID(c.Request().Context(), uuid)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (handler PlantHandler) GetPlantByCommon(c echo.Context) (postgres.Plant, error) {
	common := c.QueryParam("common")
	if common == "" {
		return postgres.Plant{}, errors.New("missing common name")
	}

	plant, err := handler.DB.GetPlantByCommon(c.Request().Context(), common)
	if err != nil {
		return postgres.Plant{}, err
	}
	return plant, nil
}

func (handler PlantHandler) CreatePlant(c echo.Context) (map[string]model.FieldError, error) {
	params := postgres.CreatePlantParams{
		Common:         c.FormValue("common"),
		Family:         c.FormValue("family"),
		Latin:          sql.NullString{String: c.FormValue("latin"), Valid: true},
		Category:       sql.NullString{String: c.FormValue("category"), Valid: true},
		Origin:         sql.NullString{String: c.FormValue("origin"), Valid: true},
		Climate:        sql.NullString{String: c.FormValue("climate"), Valid: true},
		Tempmax:        sql.NullString{String: c.FormValue("tempmax"), Valid: true},
		Tempmin:        sql.NullString{String: c.FormValue("tempmin"), Valid: true},
		Ideallight:     sql.NullString{String: c.FormValue("ideallight"), Valid: true},
		Toleratedlight: sql.NullString{String: c.FormValue("toleratedlight"), Valid: true},
		Watering:       sql.NullString{String: c.FormValue("watering"), Valid: true},
		Insects:        sql.NullString{String: c.FormValue("insects"), Valid: true},
		Diseases:       sql.NullString{String: c.FormValue("diseases"), Valid: true},
		Soil:           sql.NullString{String: c.FormValue("soil"), Valid: true},
		Repotperiod:    sql.NullString{String: c.FormValue("repotperiod"), Valid: true},
		Use:            sql.NullString{String: c.FormValue("use"), Valid: true},
	}

	paramsSlice := reflect.ValueOf(params)
	fieldErrors := make([]model.FieldError, paramsSlice.NumField())

	for i := 0; i < paramsSlice.NumField(); i++ {
		n := paramsSlice.Type().Field(i).Name
		if paramsSlice.Field(i).Interface() == "" {
			newFieldError := model.FieldError{Field: n, Message: "Please enter a valid " + n + " name"}
			fieldErrors = append(fieldErrors, newFieldError)
		}
	}

	fieldErrorsMap := make(map[string]model.FieldError)
	for _, err := range fieldErrors {
		fieldErrorsMap[err.Field] = err
	}

	if len(fieldErrorsMap) > 1 {
		return fieldErrorsMap, errors.New("Please enter valid values for all fields")
	}

	dbPlant, err := handler.DB.GetPlantByCommon(c.Request().Context(), params.Common)
	fmt.Println(dbPlant)
	if err == nil {
		fmt.Println("Plant with that name already exists.")
		fieldErrorsMap["Common"] = model.FieldError{Field: "Common", Message: "Plant with that name already exists."}
		return fieldErrorsMap, errors.New("Plant with that name already exists.")
	}

	_, createErr := handler.DB.CreatePlant(c.Request().Context(), params)
	if createErr != nil {
		return nil, createErr
	}
	return nil, nil
}

func (handler PlantHandler) DeletePlant(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	_, err = handler.DB.DeletePlant(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}
