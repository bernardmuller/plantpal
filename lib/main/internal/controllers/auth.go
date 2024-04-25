package controllers

import (
	"context"
	"database/sql"
	"domain-app/internal/store/postgres"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
)

func (controller *AuthController) GetLoginPage(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func (controller *AuthController) GetCallback(c echo.Context) error {
	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", "google"))
	googleUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		log.Println("Error while Completing Google Auth: ", err)
		return err
	}

	var newUser *postgres.User
	userIp := c.Request().RemoteAddr
	dbUser, err := controller.userService.GetUserByEmail(c.Request().Context(), googleUser.Email)
	if err != nil {
		userId := uuid.New()
		newDbUser, err := controller.userService.CreateUser(c.Request().Context(), postgres.CreateUserParams{
			ID:        userId,
			Email:     googleUser.Email,
			Firstname: googleUser.FirstName,
			Lastname:  googleUser.LastName,
			Provider:  sql.NullString{String: "Google", Valid: true},
			Image:     sql.NullString{String: googleUser.AvatarURL, Valid: true},
		})
		if err != nil {
			log.Println("Error creating new googleUser: ", err)
			return err
		}
		newUser = &newDbUser
		sessionId := uuid.New()
		newSession, sessionErr := controller.authService.CreateSession(c.Request().Context(), postgres.CreateSessionParams{
			ID:          sessionId,
			UserID:      newUser.ID,
			Expires:     googleUser.ExpiresAt,
			IpAddress:   userIp,
			AccessToken: googleUser.AccessToken,
		})
		if sessionErr != nil {
			log.Println("Error creating new session: ", err)
			return err
		}
		http.SetCookie(c.Response(), &http.Cookie{
			Name:     "plant_session",
			Value:    newSession.ID.String(),
			Secure:   false,
			HttpOnly: true,
			Expires:  googleUser.ExpiresAt,
			Path:     "/",
		})
		http.Redirect(c.Response(), req, "/plants", http.StatusFound)
	}

	sessionId := uuid.New()

	newSession, sessionErr := controller.authService.CreateSession(c.Request().Context(), postgres.CreateSessionParams{
		ID:          sessionId,
		UserID:      dbUser.ID,
		Expires:     googleUser.ExpiresAt,
		IpAddress:   userIp,
		AccessToken: googleUser.AccessToken,
	})
	if sessionErr != nil {
		log.Fatal("Error creating new session: ", err)
		return err
	}

	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "plant_session",
		Value:    newSession.ID.String(),
		Expires:  googleUser.ExpiresAt,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(c.Response(), req, "/plants", http.StatusFound)
	return nil
}

func (controller *AuthController) Logout(c echo.Context) error {
	fmt.Println(c.Request())
	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

	// get the cookie
	cookies := c.Request().Cookies()
	for _, cookie := range cookies {
		fmt.Println("Cookie: ", cookie.Name, cookie.Value)
	}
	cookie, err := c.Request().Cookie("plant_session")
	if err != nil {
		fmt.Println("Error getting cookie: ", err)
		return err
	}
	sessionId, err := uuid.Parse(cookie.Value)
	if err != nil {
		fmt.Println("Error session id: ", err)
		return err
	}

	err = controller.authService.DeleteSessionById(c.Request().Context(), sessionId)
	if err != nil {
		fmt.Println("Error deleting session: ", err)
		return err
	}

	gothic.Logout(c.Response(), c.Request())

	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "plant_session",
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	})

	http.Redirect(c.Response(), c.Request(), "/auth/login", http.StatusPermanentRedirect)
	return nil
}

func (controller *AuthController) GetProvider(c echo.Context) error {
	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		fmt.Println("Get Provider => ", gothUser)
		return nil
	} else {
		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}
