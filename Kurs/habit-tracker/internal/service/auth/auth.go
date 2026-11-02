package auth

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"habit-tracker/internal/models"
	userrepo "habit-tracker/internal/repository/user"
)

type Service struct {
	users     *userrepo.Repository
	jwtSecret []byte
	jwtTTL    time.Duration
}

func NewService(users *userrepo.Repository, jwtSecret string, jwtTTL time.Duration) *Service {
	return &Service{
		users:     users,
		jwtSecret: []byte(jwtSecret),
		jwtTTL:    jwtTTL,
	}
}

func (s *Service) Register(ctx context.Context, in models.RegisterInput) (*models.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         in.Name,
		Email:        in.Email,
		PasswordHash: string(hash),
	}
	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Login(ctx context.Context, in models.LoginInput) (string, error) {
	if err := in.Validate(); err != nil {
		return "", err
	}

	user, err := s.users.GetByEmail(ctx, in.Email)
	if errors.Is(err, models.ErrNotFound) {
		return "", models.ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)); err != nil {
		return "", models.ErrInvalidCredentials
	}
	return s.createToken(user.ID)
}

func (s *Service) Authenticate(tokenString string) (int64, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, models.ErrInvalidToken
		}
		return s.jwtSecret, nil
	})
	if err != nil || token == nil || !token.Valid {
		return 0, models.ErrInvalidToken
	}

	id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil || id <= 0 {
		return 0, models.ErrInvalidToken
	}
	return id, nil
}

func (s *Service) createToken(userID int64) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(s.jwtTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
