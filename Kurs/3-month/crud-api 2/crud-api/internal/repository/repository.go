package repository

import (
	"context"
	"crud-api/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post models.Post) (int64, error)
}
