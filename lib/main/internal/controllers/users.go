package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (controller *UsersController) GetUserById(c echo.Context) error {
	userId := c.Param("id")
	parsedId, err := uuid.Parse(userId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	user, err := controller.userService.GetUserById(c.Request().Context(), parsedId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching user")
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UsersController) GetUserBySessionId(c echo.Context) error {
	//cc := c.(*utils.CustomContext)
	fmt.Println("getUserBySessionId")
	//cookie, err := c.Cookie("plant_session")
	//if err != nil {
	//	return c.String(http.StatusUnauthorized, "Plant session cookie not found")
	//}
	//return c.String(http.StatusOK, "read a cookie")

	//if err != nil {
	//return c.String(http.StatusUnauthorized, "Unauthorized")
	//}
	//parsedCookie, _ := uuid.Parse(cookie.Value)
	//session, err := controller.authService.GetSessionById(c.Request().Context(), parsedCookie)
	//if err != nil {
	//	return c.String(http.StatusUnauthorized, "Unauthorized")
	//}
	//user, err := controller.userService.GetUserById(c.Request().Context(), session.UserID)
	//if err != nil {
	//	return c.String(http.StatusNotFound, "User not found")
	//}
	//
	return nil
	//return c.JSON(http.StatusOK, {})
}
