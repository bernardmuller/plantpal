package utils

import (
	"context"
	"domain-app/internal/views"
)

type CustomContext struct {
	context.Context
	Data     interface{}
	Renderer *views.Templates
}

func (c CustomContext) SetData(data interface{}) {
	//c("FormData", data)
	c.Data = data
}

func (c CustomContext) SetRenderer(t *views.Templates) {
	c.Renderer = t
}
