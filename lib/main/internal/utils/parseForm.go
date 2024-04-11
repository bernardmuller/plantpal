package utils

import (
	"database/sql"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func getAllFormValues(r *http.Request) map[string][]string {
	r.ParseForm()
	formValues := make(map[string][]string)

	for key, values := range r.Form {
		formValues[key] = values
	}
	return formValues
}

func PopulateStructFromForm(r *http.Request, target interface{}) interface{} {
	formValues := getAllFormValues(r)

	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)

	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		panic("target must be a pointer to a struct")
	}
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)

		fieldName := field.Name

		if fieldName == "ID" {
			continue
		}

		fieldValue := targetValue.Elem().FieldByName(fieldName)
		lowercaseFieldName := strings.ToLower(fieldName)

		if formValues[lowercaseFieldName] != nil {
			if fieldValue.Type() == reflect.TypeOf(sql.NullString{}) {
				fieldValue.Set(reflect.ValueOf(sql.NullString{String: formValues[lowercaseFieldName][0], Valid: true}))
			} else {
				if fieldValue.Kind() == reflect.Slice {
					fieldType := field.Type.Elem()
					slice := reflect.MakeSlice(fieldValue.Type(), len(formValues[lowercaseFieldName]), len(formValues[lowercaseFieldName]))

					for j, value := range formValues[lowercaseFieldName] {
						parsedValue := reflect.New(fieldType).Elem()
						parseBasicType(value, parsedValue)
						slice.Index(j).Set(parsedValue)
					}

					fieldValue.Set(slice)
				} else {
					parseBasicType(formValues[lowercaseFieldName][0], fieldValue)
				}
			}
		}
	}

	return target
}

func parseBasicType(value string, fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			fieldValue.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			fieldValue.SetUint(uintValue)
		}
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err == nil {
			fieldValue.SetFloat(floatValue)
		}
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			fieldValue.SetBool(boolValue)
		}
	}
}
