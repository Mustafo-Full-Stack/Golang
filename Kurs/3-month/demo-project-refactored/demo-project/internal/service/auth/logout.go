package auth

import (
	"context"

	"demo-project/internal/models"
)

// Logout завершает текущую сессию пользователя.
func (s *Service) Logout(ctx context.Context, user models.CurrentUser) error {
	return s.RevokeSession(ctx, user, user.SessionID)
}
