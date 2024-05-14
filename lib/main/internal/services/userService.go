package services

import (
	"context"
	"domain-app/internal/store/postgres"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IUserService interface {
	GetUserByEmail(c echo.Context, email string) (postgres.User, error)
	//CreateUserSession(c echo.Context, access_token string) ([]postgres.Plant, error)
	GetUserById(c echo.Context, id string) (postgres.User, error)
}

type UserDBService struct {
	DB *postgres.Queries
}

func (service UserDBService) GetUserByEmail(c context.Context, email string) (postgres.User, error) {
	user, err := service.DB.GetUserByEmail(c, email)
	if err != nil {
		return postgres.User{}, err
	}
	return user, nil
}

func (service UserDBService) CreateUser(c context.Context, params postgres.CreateUserParams) (postgres.User, error) {
	newUser, createErr := service.DB.CreateUser(c, params)
	if createErr != nil {
		return postgres.User{}, createErr
	}
	return newUser, nil
}

func (service UserDBService) GetUserById(c context.Context, id uuid.UUID) (postgres.User, error) {
	user, err := service.DB.GetUserById(c, id)
	if err != nil {
		return postgres.User{}, err
	}
	return user, nil
}
