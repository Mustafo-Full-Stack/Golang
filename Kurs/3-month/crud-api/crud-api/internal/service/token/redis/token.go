package redis

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/service/token"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	goredis "github.com/redis/go-redis/v9"
)

const keyPrefix = "token:"

var _ token.Manager = (*tokenManager)(nil)

type tokenManager struct {
	client *goredis.Client
	ttl    time.Duration
}

func NewTokenManager(client *goredis.Client, ttl time.Duration) token.Manager {
	return &tokenManager{client: client, ttl: ttl}
}

func (m *tokenManager) Generate(ctx context.Context, userID models.ID) (string, error) {
	value := uuid.NewString()

	if err := m.client.Set(ctx, keyPrefix+value, int64(userID), m.ttl).Err(); err != nil {
		return "", fmt.Errorf("save token: %w", err)
	}

	return value, nil
}

func (m *tokenManager) Validate(ctx context.Context, value string) (models.ID, error) {
	raw, err := m.client.Get(ctx, keyPrefix+value).Result()
	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return 0, models.ErrInvalidToken
		}
		return 0, fmt.Errorf("get token: %w", err)
	}

	userID, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse token user id: %w", err)
	}

	return models.ID(userID), nil
}

func (m *tokenManager) Revoke(ctx context.Context, value string) error {
	deleted, err := m.client.Del(ctx, keyPrefix+value).Result()
	if err != nil {
		return fmt.Errorf("delete token: %w", err)
	}
	if deleted == 0 {
		return models.ErrInvalidToken
	}

	return nil
}
