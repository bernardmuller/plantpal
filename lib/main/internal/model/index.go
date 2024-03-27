package model

import "domain-app/internal/store/postgres"

type Data struct {
	Plants []postgres.Plant
}

func NewPageData(data Data, form FormData) PageData {
	return PageData{
		Data: data,
		Form: form,
	}
}

type FormData struct {
	Errors      map[string]string
	FieldErrors map[string]FieldError
	Values      map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors:      map[string]string{},
		FieldErrors: map[string]FieldError{},
		Values:      map[string]string{},
	}
}

type PageData struct {
	Data Data
	Form FormData
}
