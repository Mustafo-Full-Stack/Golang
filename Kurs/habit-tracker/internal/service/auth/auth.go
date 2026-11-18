package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"habit-tracker/internal/models"
	userrepo "habit-tracker/internal/repository/user"
)

const (
	accessTokenType  = "access"
	refreshTokenType = "refresh"
)

type tokenClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

type Service struct {
	users      *userrepo.Repository
	jwtSecret  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewService(users *userrepo.Repository, jwtSecret string, accessTTL, refreshTTL time.Duration) *Service {
	return &Service{
		users:      users,
		jwtSecret:  []byte(jwtSecret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
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

func (s *Service) Login(ctx context.Context, in models.LoginInput) (models.TokenPair, error) {

	if err := in.Validate(); err != nil {
		return models.TokenPair{}, err
	}

	user, err := s.users.GetByEmail(ctx, in.Email)

	if errors.Is(err, models.ErrNotFound) {
		return models.TokenPair{}, models.ErrInvalidCredentials
	}

	if err != nil {
		return models.TokenPair{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)); err != nil {
		return models.TokenPair{}, models.ErrInvalidCredentials
	}

	accessToken, err := s.createToken(user.ID, accessTokenType, s.accessTTL)
	if err != nil {
		return models.TokenPair{}, err
	}

	refreshToken, err := s.createToken(user.ID, refreshTokenType, s.refreshTTL)
	if err != nil {
		return models.TokenPair{}, err
	}

	return models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Authenticate(tokenString string) (int64, error) {
	claims := &tokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, models.ErrInvalidToken
		}
		return s.jwtSecret, nil
	})

	if err != nil || token == nil || !token.Valid {
		return 0, models.ErrInvalidToken
	}

	if claims.Type != accessTokenType || claims.UserID <= 0 {
		return 0, models.ErrInvalidToken
	}

	return claims.UserID, nil
}

func (s *Service) createToken(userID int64, tokenType string, ttl time.Duration) (string, error) {
	now := time.Now()

	claims := tokenClaims{
		UserID: userID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtSecret)
}
