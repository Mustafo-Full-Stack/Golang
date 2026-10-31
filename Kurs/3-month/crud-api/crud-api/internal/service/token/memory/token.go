package memory

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/service/token"
	"sync"
	"time"

	"github.com/google/uuid"
)

var _ token.Manager = (*tokenManager)(nil)

type entry struct {
	userID    models.ID
	expiresAt time.Time
}

type tokenManager struct {
	mu     sync.Mutex
	ttl    time.Duration
	tokens map[string]entry
}

func NewTokenManager(ttl time.Duration) token.Manager {
	return &tokenManager{
		ttl:    ttl,
		tokens: make(map[string]entry),
	}
}

func (m *tokenManager) Generate(ctx context.Context, userID models.ID) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	value := uuid.NewString()

	m.mu.Lock()
	defer m.mu.Unlock()

	m.tokens[value] = entry{userID: userID, expiresAt: time.Now().Add(m.ttl)}

	return value, nil
}

func (m *tokenManager) Validate(ctx context.Context, value string) (models.ID, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	e, ok := m.lookup(value)
	if !ok {
		return 0, models.ErrInvalidToken
	}

	return e.userID, nil
}

func (m *tokenManager) Revoke(ctx context.Context, value string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.lookup(value); !ok {
		return models.ErrInvalidToken
	}

	delete(m.tokens, value)
	return nil
}

// lookup лениво удаляет просроченный токен, чтобы map не копила мусор.
func (m *tokenManager) lookup(value string) (entry, bool) {
	e, ok := m.tokens[value]
	if !ok {
		return entry{}, false
	}
	if time.Now().After(e.expiresAt) {
		delete(m.tokens, value)
		return entry{}, false
	}

	return e, true
}
