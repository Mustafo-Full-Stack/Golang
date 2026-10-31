package auth

import (
	"time"

	"demo-project/internal/models"
	"demo-project/internal/repository"
	"demo-project/internal/service"
)

var _ service.AuthService = (*Service)(nil)

// TokenManager — подпись и разбор JWT. Интерфейс позволяет подставить фейк в тестах.
type TokenManager interface {
	GenerateAccess(userID, roleID int64, sessionID string) (string, error)
	GenerateRefresh(userID int64, sessionID string) (string, error)
	ParseAccess(token string) (models.CurrentUser, error)
	ParseRefresh(token string) (userID int64, sessionID string, err error)
	AccessTTL() time.Duration
	RefreshTTL() time.Duration
}

type Service struct {
	users    repository.UserRepository
	sessions repository.SessionRepository
	tokens   repository.TokenRepository
	jwt      TokenManager
}

func NewService(
	users repository.UserRepository,
	sessions repository.SessionRepository,
	tokens repository.TokenRepository,
	jwt TokenManager,
) *Service {
	return &Service{
		users:    users,
		sessions: sessions,
		tokens:   tokens,
		jwt:      jwt,
	}
}
