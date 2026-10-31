package auth

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

func (s *Service) ListSessions(ctx context.Context, user models.CurrentUser) ([]models.Session, error) {
	sessions, err := s.sessions.List(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("list sessions: %w", err)
	}
	return sessions, nil
}
