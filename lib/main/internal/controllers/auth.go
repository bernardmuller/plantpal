package controllers

import (
	"database/sql"
	"domain-app/internal/store/postgres"
	"encoding/json"
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
	//cookie, err := c.Request().Cookie("plant_session")
	//if err != nil {
	//	http.Redirect(c.Response(), c.Request(), "/auth/login", http.StatusFound)
	//}
	//
	//sessionId, err := uuid.Parse(cookie.Value)
	//if err != nil {
	//	http.Redirect(c.Response(), c.Request(), "/auth/login", http.StatusFound)
	//}
	//
	//_, err = controller.authService.GetSessionById(c.Request().Context(), sessionId)
	//if err == nil {
	//	http.Redirect(c.Response(), c.Request(), "/plants", http.StatusFound)
	//}

	//req := c.Request().WithContext(context.WithValue(context.Background(), "provider", "google"))
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
		json.NewEncoder(c.Response()).Encode(newUser)
	}

	newSessionId := uuid.New()

	newSession, sessionErr := controller.authService.CreateSession(c.Request().Context(), postgres.CreateSessionParams{
		ID:          newSessionId,
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

	//json.NewEncoder(c.Response()).Encode(dbUser)
	http.Redirect(c.Response(), c.Request(), "http://localhost:5173/plants", http.StatusTemporaryRedirect)
	return nil
}

func (controller *AuthController) Logout(c echo.Context) error {
	fmt.Println(c.Request())
	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

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

	json.NewEncoder(c.Response()).Encode("Logged out")
	return nil
}

func (controller *AuthController) GetProvider(c echo.Context) error {
	fmt.Println("Get Provider => ", c.Request().URL.Query().Get("provider"))

	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		fmt.Println("User => ", gothUser)
		return nil
	} else {
		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}
