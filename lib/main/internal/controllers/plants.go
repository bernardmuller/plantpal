package controllers

import (
	"database/sql"
	"domain-app/internal/config"
	"domain-app/internal/model"
	"domain-app/internal/services"
	"domain-app/internal/store/postgres"
	"domain-app/internal/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type PlantsController struct {
	config *config.ApiConfig
}

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

func (controller *PlantsController) GetAllPlants(w http.ResponseWriter, r *http.Request) {
	// Retrieve plants from the database
	plants, err := services.PlantsDbService{DB: controller.config.Database}.GetAllPlants(r.Context())
	if err != nil {
		http.Error(w, "Error fetching plants", http.StatusInternalServerError)
	}

	// Construct page data
	pageData := model.Data{Plants: plants}

	// Render the response
	c, ok := r.Context().Value("customContext").(utils.CustomContext)
	if !ok {
		// Handle the case where custom context is not found or of unexpected type
		fmt.Println("Custom context not found")
		http.Error(w, "Custom context not found", http.StatusInternalServerError)
		return
	}
	c.Renderer.Render(w, "plants", pageData)
}

func (controller *PlantsController) GetCreatePlantForm(c echo.Context) error {
	formData := model.NewFormData()
	return c.Render(200, "createPlant", formData)
}
func ValidateFormData(formData interface{}) (map[string]model.FieldError, error) {
	fmt.Println(formData)
	paramsSlice := reflect.ValueOf(formData)

	if paramsSlice.Kind() != reflect.Ptr {
		ptr := reflect.New(reflect.TypeOf(formData))
		ptr.Elem().Set(paramsSlice)
		paramsSlice = ptr
	}

	paramsSlice = paramsSlice.Elem()

	fieldErrors := make([]model.FieldError, 0)

	for i := 0; i < paramsSlice.NumField(); i++ {
		field := paramsSlice.Type().Field(i)
		fieldValue := paramsSlice.Field(i)

		if field.Name == "ID" {
			continue
		}

		if fieldValue.Type() == reflect.TypeOf(sql.NullString{}) {
			continue
		} else {
			if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(field.Type).Interface()) {
				fieldErrors = append(fieldErrors, model.FieldError{
					Field:   field.Name,
					Message: "Please enter a valid " + field.Name + " name",
				})
			}
		}
	}

	// Convert fieldErrors slice to a map
	fieldErrorsMap := make(map[string]model.FieldError)
	for _, err := range fieldErrors {
		fieldErrorsMap[err.Field] = err
	}

	if len(fieldErrors) > 0 {
		return fieldErrorsMap, errors.New("Please enter valid values for all fields")
	}

	return fieldErrorsMap, nil
}

func (controller *PlantsController) CreatePlant(c echo.Context) error {
	var target postgres.CreatePlantParams
	var err string
	formData := utils.PopulateStructFromForm(c.Request(), &target).(*postgres.CreatePlantParams)
	var fieldErrors = make(map[string]model.FieldError)
	validationErrors, valErr := ValidateFormData(formData)
	if valErr != nil {
		err = valErr.Error()
		for i, j := range validationErrors {
			fieldErrors[i] = j
		}
	}

	_, lookupErr := services.PlantsDbService{
		DB: controller.config.Database,
	}.GetPlantByCommon(c.Request().Context(), formData.Common)
	if lookupErr == nil {
		errMessage := "Plant with that name already exists."
		fieldErrors["Common"] = model.FieldError{Field: "Common", Message: errMessage}
		err = errMessage
	}

	if len(fieldErrors) == 0 {
		var data postgres.CreatePlantParams
		formData.ID = uuid.New()
		data = *formData
		p, createErr := services.PlantsDbService{
			DB: controller.config.Database,
		}.CreatePlant(c.Request().Context(), data)

		if createErr != nil {
			err = createErr.Error()
		}
		fmt.Println(p)
	}

	if err != "" || len(fieldErrors) != 0 {
		formData := model.FormData{
			Errors: map[string]string{
				"error": err,
			},
			FieldErrors: fieldErrors,
			Values: map[string]string{
				"common":         c.FormValue("common"),
				"family":         c.FormValue("family"),
				"latin":          c.FormValue("latin"),
				"category":       c.FormValue("category"),
				"origin":         c.FormValue("origin"),
				"climate":        c.FormValue("climate"),
				"tempmax":        c.FormValue("tempmax"),
				"tempmin":        c.FormValue("tempmin"),
				"ideallight":     c.FormValue("ideallight"),
				"toleratedlight": c.FormValue("toleratedlight"),
				"watering":       c.FormValue("watering"),
				"insects":        c.FormValue("insects"),
				"diseases":       c.FormValue("diseases"),
				"soil":           c.FormValue("soil"),
				"repotperiod":    c.FormValue("repotperiod"),
				"use":            c.FormValue("use"),
			}}
		return c.Render(422, "createPlantForm", formData)
	}

	return c.Redirect(303, "/plants")
}

func (controller *PlantsController) DeletePlant(c echo.Context) error {
	id := c.Param("id")
	parsedId, parseErr := uuid.Parse(id)
	if parseErr != nil {
		return parseErr
	}
	err := services.PlantsDbService{DB: controller.config.Database}.DeletePlant(c.Request().Context(), parsedId)
	if err != nil {
		return c.String(400, err.Error())
	}

	return c.NoContent(200)
}
