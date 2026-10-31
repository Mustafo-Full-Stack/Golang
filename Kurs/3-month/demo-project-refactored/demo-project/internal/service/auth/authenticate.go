package auth

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

// Authenticate проверяет access-токен: подпись и срок (JWT), а затем — что это
// текущий токен сессии в Redis. После logout или refresh старый токен не пройдёт.
func (s *Service) Authenticate(ctx context.Context, accessToken string) (models.CurrentUser, error) {
	user, err := s.jwt.ParseAccess(accessToken)
	if err != nil {
		return models.CurrentUser{}, err
	}

	ok, err := s.tokens.IsAccessValid(ctx, user.SessionID, accessToken)
	if err != nil {
		return models.CurrentUser{}, fmt.Errorf("check access token: %w", err)
	}
	if !ok {
		return models.CurrentUser{}, models.ErrInvalidToken
	}
	return user, nil
}
