package repository

import (
	"context"
	"crud-api/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (models.ID, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id models.ID) (*models.User, error)
}
