package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"demo-project/internal/models"
)

func (s *Service) Register(ctx context.Context, in models.RegisterInput) (*models.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	u := &models.User{
		Email:        in.Email,
		PasswordHash: string(hash),
		Name:         in.Name,
		RoleID:       models.RoleUser,
	}
	// Занятость email не проверяем заранее: это сделает UNIQUE в БД (models.ErrEmailTaken).
	if err := s.users.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}
