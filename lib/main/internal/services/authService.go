package services

import (
	"context"
	"domain-app/internal/store/postgres"
	"github.com/google/uuid"
)

type IAuthService interface {
	CreateSession(c context.Context, params postgres.CreateSessionParams) (postgres.Session, error)
	GetSessionById(c context.Context, id uuid.UUID) (postgres.Session, error)
	DeleteSessionById(c context.Context, id uuid.UUID) error
}

type AuthDBService struct {
	DB *postgres.Queries
}

func (service AuthDBService) CreateSession(c context.Context, params postgres.CreateSessionParams) (postgres.Session, error) {
	session, err := service.DB.CreateSession(c, params)
	if err != nil {
		return postgres.Session{}, err
	}
	return session, nil
}

func (service AuthDBService) GetSessionById(c context.Context, id uuid.UUID) (postgres.Session, error) {
	session, err := service.DB.GetSessionByID(c, id)
	if err != nil {
		return postgres.Session{}, err
	}
	return session, nil
}

func (service AuthDBService) DeleteSessionById(c context.Context, id uuid.UUID) error {
	err := service.DB.DeleteSessionByID(c, id)
	if err != nil {
		return err
	}
	return nil
}
