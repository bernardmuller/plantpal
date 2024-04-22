package controllers

import (
	"context"
	"domain-app/internal/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"net/http"
)

type AuthController struct {
	config *config.ApiConfig
}

func (controller *AuthController) GetLoginPage(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func (controller *AuthController) GetCallback(c echo.Context) error {
	//provider := c.Param("provider")
	fmt.Println("Callback")
	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", "google"))
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		fmt.Fprintln(c.Response(), err)
		return err
	}
	fmt.Println("Authenticated user: ", user)
	http.Redirect(c.Response(), req, "/plants", http.StatusFound)
	return nil
}

func (controller *AuthController) Logout(c echo.Context) error {
	gothic.Logout(c.Response(), c.Request())
	c.Response().Header().Set("Location", "/")
	c.Response().WriteHeader(http.StatusTemporaryRedirect)
	return nil
}

func (controller *AuthController) GetProvider(c echo.Context) error {
	// try to get the user without re-authenticating

	fmt.Println("Get Provider =======================================")
	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		//t, _ := template.New("foo").Parse(userTemplate)
		//t.Execute(res, gothUser)
		fmt.Println("Get Provider => ", gothUser)
		return nil
	} else {
		fmt.Println("begin auth handler")
		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}
