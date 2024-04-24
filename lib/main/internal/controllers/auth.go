package controllers

import (
	"context"
	"database/sql"
	"domain-app/internal/config"
	"domain-app/internal/services"
	"domain-app/internal/store/postgres"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
)

type AuthController struct {
	config *config.ApiConfig
}

func (controller *AuthController) GetLoginPage(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func (controller *AuthController) GetCallback(c echo.Context) error {
	req := c.Request().WithContext(context.WithValue(context.Background(), "provider", "google"))
	googleUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		fmt.Println("Error while Completing Google Auth: ", err)
		return err
	}
	fmt.Println("Authenticated googleUser: ", googleUser)

	var newUser *postgres.User
	fmt.Println("testtest")
	userService := services.UserDBService{DB: controller.config.Database}

	fmt.Println("testtest12")
	userIp := c.Request().RemoteAddr
	fmt.Println("test1")
	dbUser, err := userService.GetUserByEmail(c.Request().Context(), googleUser.Email)
	if err != nil {
		fmt.Println("test2")
		userId := uuid.New()
		newDbUser, err := userService.CreateUser(c.Request().Context(), postgres.CreateUserParams{
			ID:        userId,
			Email:     googleUser.Email,
			Firstname: googleUser.FirstName,
			Lastname:  googleUser.LastName,
			Provider:  sql.NullString{String: "Google", Valid: true},
			Image:     sql.NullString{String: googleUser.AvatarURL, Valid: true},
		})
		if err != nil {
			log.Fatal("Error creating new googleUser: ", err)
			return err
		}
		newUser = &newDbUser
		fmt.Println("test2a")
		sessionId := uuid.New()
		newSession, sessionErr := services.AuthDBService{DB: controller.config.Database}.CreateSession(c.Request().Context(), postgres.CreateSessionParams{
			ID:          sessionId,
			UserID:      newUser.ID,
			Expires:     googleUser.ExpiresAt,
			IpAddress:   userIp,
			AccessToken: googleUser.AccessToken,
		})
		if sessionErr != nil {
			log.Fatal("Error creating new session: ", err)
			return err
		}
		fmt.Println("test2b")
		http.SetCookie(c.Response(), &http.Cookie{
			Name:  "plant_session",
			Value: newSession.ID.String(),
		})
		http.Redirect(c.Response(), req, "/plants", http.StatusFound)
	}

	sessionId := uuid.New()

	newSession, sessionErr := services.AuthDBService{DB: controller.config.Database}.CreateSession(c.Request().Context(), postgres.CreateSessionParams{
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
		Name:    "plant_session",
		Value:   newSession.ID.String(),
		Expires: googleUser.ExpiresAt,
	})
	http.Redirect(c.Response(), req, "/plants", http.StatusFound)
	return nil
}

func (controller *AuthController) Logout(c echo.Context) error {
	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

	// get the cookie
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

	err = services.AuthDBService{DB: controller.config.Database}.DeleteSessionById(c.Request().Context(), sessionId)
	if err != nil {
		fmt.Println("Error deleting session: ", err)
		return err
	}

	gothic.Logout(c.Response(), c.Request())

	c.Response().Header().Set("Location", "/auth/login")
	c.Response().WriteHeader(http.StatusTemporaryRedirect)
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
