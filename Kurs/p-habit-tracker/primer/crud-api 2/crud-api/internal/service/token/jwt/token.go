package jwt

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/service/token"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenType  = "access"
	refreshTokenType = "refresh"
)

var _ token.Manager = (*tokenManager)(nil)

type claims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

type tokenManager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewTokenManager(secret string, accessTTL, refreshTTL time.Duration) token.Manager {
	return &tokenManager{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (m *tokenManager) GenerateAccessToken(_ context.Context, userID int64) (string, error) {
	return m.generate(userID, accessTokenType, m.accessTTL)
}

func (m *tokenManager) GenerateRefreshToken(_ context.Context, userID int64) (string, error) {
	return m.generate(userID, refreshTokenType, m.refreshTTL)
}

func (m *tokenManager) Validate(_ context.Context, value string) (int64, error) {
	var c claims

	// ParseWithClaims проверяет подпись и срок жизни и заполняет c данными из токена.
	_, err := jwt.ParseWithClaims(value, &c, m.secretKey,
		// Разрешаем только HS256, чтобы нельзя было подсунуть токен с другим алгоритмом.
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return 0, models.ErrInvalidToken
	}

	if c.Type != accessTokenType {
		return 0, models.ErrInvalidToken
	}

	return c.UserID, nil
}

func (m *tokenManager) generate(userID int64, tokenType string, ttl time.Duration) (string, error) {
	now := time.Now()

	c := claims{
		UserID:    userID,
		Type:      tokenType,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	signed, err := jwtToken.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("sign %s token: %w", tokenType, err)
	}

	return signed, nil
}

func (m *tokenManager) secretKey(*jwt.Token) (any, error) {
	return m.secret, nil
}
