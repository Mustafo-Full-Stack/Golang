package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"demo-project/internal/models"
)

// RevokeSession удаляет сессию пользователя и её токены. Чужую сессию удалить нельзя:
// репозиторий ищет её по паре (userID, sessionID) и вернёт ErrNotFound.
func (s *Service) RevokeSession(ctx context.Context, user models.CurrentUser, sessionID string) error {
	if err := uuid.Validate(sessionID); err != nil {
		return fmt.Errorf("%w: invalid session id", models.ErrInvalidInput)
	}

	if err := s.sessions.Delete(ctx, user.ID, sessionID); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	if err := s.tokens.Delete(ctx, sessionID); err != nil {
		return fmt.Errorf("delete session tokens: %w", err)
	}
	return nil
}

// RevokeOtherSessions удаляет все сессии пользователя, кроме текущей.
func (s *Service) RevokeOtherSessions(ctx context.Context, user models.CurrentUser) error {
	ids, err := s.sessions.DeleteOthers(ctx, user.ID, user.SessionID)
	if err != nil {
		return fmt.Errorf("delete other sessions: %w", err)
	}
	if err := s.tokens.DeleteMany(ctx, ids); err != nil {
		return fmt.Errorf("delete session tokens: %w", err)
	}
	return nil
}
