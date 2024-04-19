package utils

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
	Data        interface{}
	CurrentUser interface{}
}

func (c CustomContext) SetData(data interface{}) {
	c.Set("FormData", data)
}
