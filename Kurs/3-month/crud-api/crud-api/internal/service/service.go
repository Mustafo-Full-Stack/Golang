package service

import (
	"context"
	"crud-api/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, name, email, password string) (models.ID, error)
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, token string) error
	Authenticate(ctx context.Context, token string) (models.ID, error)
}

type UserService interface {
	GetByID(ctx context.Context, id models.ID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type Services struct {
	AuthService AuthService
	UserService UserService
}
