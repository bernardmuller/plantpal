package controllers

import (
	"domain-app/internal/config"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	config *config.ApiConfig
}

func (controller *AuthController) GetLoginPage(c echo.Context) error {
	return c.Render(200, "login", nil)
}
