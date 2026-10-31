package token

import (
	"context"
	"crud-api/internal/models"
)

type Manager interface {
	Generate(ctx context.Context, userID models.ID) (string, error)
	Validate(ctx context.Context, token string) (models.ID, error)
	Revoke(ctx context.Context, token string) error
}
