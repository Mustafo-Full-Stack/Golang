package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"demo-project/internal/models"
)

// Login создаёт сессию и выдаёт пару токенов.
// На неизвестный email и неверный пароль ответ одинаковый: не подсказываем, что именно неверно.
func (s *Service) Login(ctx context.Context, in models.LoginInput) (*models.TokenPair, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	u, err := s.users.GetByEmail(ctx, in.Email)
	if errors.Is(err, models.ErrNotFound) {
		return nil, models.ErrInvalidCredentials
	}
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
		return nil, models.ErrInvalidCredentials
	}

	session := &models.Session{
		ID:         uuid.NewString(),
		UserID:     u.ID,
		DeviceType: in.DeviceType,
		IP:         in.IP,
		UserAgent:  in.UserAgent,
		ExpiresAt:  time.Now().Add(s.jwt.RefreshTTL()),
	}
	if err := s.sessions.Create(ctx, session); err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	pair, err := s.issueTokens(ctx, u, session.ID)
	if err != nil {
		// Сессия без токенов не нужна: удаляем её, не теряя ни одной ошибки.
		if deleteErr := s.sessions.Delete(ctx, u.ID, session.ID); deleteErr != nil {
			return nil, errors.Join(err, fmt.Errorf("delete session: %w", deleteErr))
		}
		return nil, err
	}
	return pair, nil
}

// issueTokens подписывает новую пару токенов для сессии и сохраняет её в Redis.
func (s *Service) issueTokens(ctx context.Context, u *models.User, sessionID string) (*models.TokenPair, error) {
	access, err := s.jwt.GenerateAccess(u.ID, u.RoleID, sessionID)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}
	refresh, err := s.jwt.GenerateRefresh(u.ID, sessionID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	pair := &models.TokenPair{AccessToken: access, RefreshToken: refresh}
	if err := s.tokens.SavePair(ctx, sessionID, pair, s.jwt.AccessTTL(), s.jwt.RefreshTTL()); err != nil {
		return nil, fmt.Errorf("save tokens: %w", err)
	}
	return pair, nil
}
