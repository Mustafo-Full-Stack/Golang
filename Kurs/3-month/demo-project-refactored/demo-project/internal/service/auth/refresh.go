package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"demo-project/internal/models"
)

// Refresh выдаёт новую пару токенов для той же сессии.
// Старый refresh-токен одноразовый: UseRefresh удаляет его из Redis.
func (s *Service) Refresh(ctx context.Context, refreshToken string) (*models.TokenPair, error) {
	userID, sessionID, err := s.jwt.ParseRefresh(refreshToken)
	if err != nil {
		return nil, err
	}

	// Роль читаем заново: в refresh-токене её нет, и она могла измениться.
	// Читаем до UseRefresh, чтобы сбой БД не «сжёг» токен впустую.
	u, err := s.users.GetByID(ctx, userID)
	if errors.Is(err, models.ErrNotFound) {
		return nil, models.ErrInvalidToken
	}
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	ok, err := s.tokens.UseRefresh(ctx, sessionID, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("use refresh token: %w", err)
	}
	if !ok {
		return nil, models.ErrInvalidToken
	}

	err = s.sessions.Extend(ctx, sessionID, time.Now().Add(s.jwt.RefreshTTL()))
	if errors.Is(err, models.ErrNotFound) {
		return nil, models.ErrInvalidToken
	}
	if err != nil {
		return nil, fmt.Errorf("extend session: %w", err)
	}

	return s.issueTokens(ctx, u, sessionID)
}
