package user

import (
	"context"
	"fmt"

	"demo-project/internal/models"
	"demo-project/internal/repository"
	"demo-project/internal/service"
)

var _ service.UserService = (*Service)(nil)

type Service struct {
	users repository.UserRepository
}

func NewService(users repository.UserRepository) *Service {
	return &Service{users: users}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.User, error) {
	u, err := s.users.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return u, nil
}
