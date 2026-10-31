package token

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"demo-project/internal/models"
)

type Repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{client: client}
}

// SavePair хранит хеши токенов, а не сами JWT: утечка Redis не даёт готовых токенов.
func (r *Repository) SavePair(ctx context.Context, sessionID string, pair *models.TokenPair, accessTTL, refreshTTL time.Duration) error {
	if err := r.client.Set(ctx, accessKey(sessionID), hashToken(pair.AccessToken), accessTTL).Err(); err != nil {
		return fmt.Errorf("save access token: %w", err)
	}
	if err := r.client.Set(ctx, refreshKey(sessionID), hashToken(pair.RefreshToken), refreshTTL).Err(); err != nil {
		return fmt.Errorf("save refresh token: %w", err)
	}
	return nil
}

// IsAccessValid проверяет, что это текущий access-токен сессии, а не старый или чужой.
func (r *Repository) IsAccessValid(ctx context.Context, sessionID, accessToken string) (bool, error) {
	stored, err := r.client.Get(ctx, accessKey(sessionID)).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("get access token: %w", err)
	}

	return stored == hashToken(accessToken), nil
}

// UseRefresh делает refresh-токен одноразовым. GETDEL читает и удаляет ключ одной командой:
// если два запроса пришли с одним токеном одновременно, ключ получит только первый.
// Старый access удаляем сразу: новая пара заменит оба токена. Если предъявили старый refresh,
// текущая пара тоже сгорает — повторное использование refresh похоже на кражу.
func (r *Repository) UseRefresh(ctx context.Context, sessionID, refreshToken string) (bool, error) {
	stored, err := r.client.GetDel(ctx, refreshKey(sessionID)).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("get refresh token: %w", err)
	}

	if err := r.client.Del(ctx, accessKey(sessionID)).Err(); err != nil {
		return false, fmt.Errorf("delete access token: %w", err)
	}
	return stored == hashToken(refreshToken), nil
}

func (r *Repository) Delete(ctx context.Context, sessionID string) error {
	if err := r.client.Del(ctx, accessKey(sessionID), refreshKey(sessionID)).Err(); err != nil {
		return fmt.Errorf("delete session tokens: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMany(ctx context.Context, sessionIDs []string) error {
	if len(sessionIDs) == 0 {
		return nil
	}

	keys := make([]string, 0, len(sessionIDs)*2)
	for _, id := range sessionIDs {
		keys = append(keys, accessKey(id), refreshKey(id))
	}
	if err := r.client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("delete session tokens: %w", err)
	}
	return nil
}

func accessKey(sessionID string) string {
	return fmt.Sprintf("session:%s:access", sessionID)
}

func refreshKey(sessionID string) string {
	return fmt.Sprintf("session:%s:refresh", sessionID)
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
