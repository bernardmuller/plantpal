package middleware

import (
	"domain-app/internal/utils"
	"github.com/labstack/echo/v4"
)

func CreateCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customContext := utils.CustomContext{
			Context: c,
			Data:    nil,
		}

		cc := customContext
		return next(cc)
	}
}
