package service

import (
	"context"
	"crud-api/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, name, email, password string) (int64, error)
	Login(ctx context.Context, email, password string) (models.TokenPair, error)
	Authenticate(ctx context.Context, token string) (int64, error)
}

type UserService interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post models.Post) (int64, error)
}

type Services struct {
	AuthService AuthService
	UserService UserService
	PostService PostService
}
