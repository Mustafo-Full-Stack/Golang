package auth

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/repository"
	"crud-api/internal/service"
	"crud-api/internal/service/token"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	maxPasswordBytes  = 72
)

var _ service.AuthService = (*authService)(nil)

type authService struct {
	users  repository.UserRepository
	tokens token.Manager
}

func NewAuthService(users repository.UserRepository, tokens token.Manager) service.AuthService {
	return &authService{
		users:  users,
		tokens: tokens,
	}
}

func (s *authService) Register(ctx context.Context, name, email, password string) (models.ID, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if name == "" || email == "" || password == "" {
		return 0, models.ErrInvalidInput
	}

	if utf8.RuneCountInString(password) < minPasswordLength {
		return 0, models.ErrPasswordTooShort
	}

	if len(password) > maxPasswordBytes {
		return 0, models.ErrPasswordTooLong
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	userID, err := s.users.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, models.ErrEmailAlreadyExists) {
			return 0, err
		}
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userID, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	if email == "" || password == "" {
		return "", models.ErrInvalidCredentials
	}

	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return "", models.ErrInvalidCredentials
		}

		return "", fmt.Errorf("failed to get user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", models.ErrInvalidCredentials
	}

	tkn, err := s.tokens.Generate(ctx, user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tkn, nil
}

func (s *authService) Logout(ctx context.Context, tokenValue string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if strings.TrimSpace(tokenValue) == "" {
		return models.ErrInvalidToken
	}

	err := s.tokens.Revoke(ctx, tokenValue)
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}

	return nil
}

func (s *authService) Authenticate(ctx context.Context, tokenValue string) (models.ID, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	if strings.TrimSpace(tokenValue) == "" {
		return 0, models.ErrInvalidToken
	}

	userID, err := s.tokens.Validate(ctx, tokenValue)
	if err != nil {
		return 0, fmt.Errorf("failed to validate token: %w", err)
	}

	return userID, nil
}
