package token

import (
	"context"
)

type Manager interface {
	GenerateAccessToken(ctx context.Context, userID int64) (string, error)
	GenerateRefreshToken(ctx context.Context, userID int64) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
}
