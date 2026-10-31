package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"demo-project/internal/models"
)

const (
	accessTokenType  = "access"
	refreshTokenType = "refresh"
)

var _ TokenManager = (*JWTManager)(nil)

type JWTConfig struct {
	Secret     []byte
	Issuer     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

// JWTManager владеет секретом, issuer и TTL. Бизнес-сервис auth про криптографию не знает.
type JWTManager struct {
	secret     []byte
	issuer     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTManager(cfg JWTConfig) (*JWTManager, error) {
	if len(cfg.Secret) < 32 {
		return nil, errors.New("jwt secret must be at least 32 bytes")
	}
	if cfg.Issuer == "" {
		return nil, errors.New("jwt issuer is required")
	}
	if cfg.AccessTTL <= 0 || cfg.RefreshTTL <= 0 {
		return nil, errors.New("jwt ttl must be positive")
	}
	if cfg.AccessTTL >= cfg.RefreshTTL {
		return nil, errors.New("jwt access ttl must be shorter than refresh ttl")
	}

	return &JWTManager{
		secret:     cfg.Secret,
		issuer:     cfg.Issuer,
		accessTTL:  cfg.AccessTTL,
		refreshTTL: cfg.RefreshTTL,
	}, nil
}

func (m *JWTManager) AccessTTL() time.Duration { return m.accessTTL }

func (m *JWTManager) RefreshTTL() time.Duration { return m.refreshTTL }

func (m *JWTManager) GenerateAccess(userID, roleID int64, sessionID string) (string, error) {
	return m.sign(tokenClaims{
		RegisteredClaims: m.registeredClaims(userID, sessionID, m.accessTTL),
		Type:             accessTokenType,
		RoleID:           roleID,
	})
}

// GenerateRefresh не кладёт роль в токен: при обновлении она читается заново из БД.
func (m *JWTManager) GenerateRefresh(userID int64, sessionID string) (string, error) {
	return m.sign(tokenClaims{
		RegisteredClaims: m.registeredClaims(userID, sessionID, m.refreshTTL),
		Type:             refreshTokenType,
	})
}

func (m *JWTManager) ParseAccess(token string) (models.CurrentUser, error) {
	claims, userID, err := m.parse(token, accessTokenType)
	if err != nil {
		return models.CurrentUser{}, err
	}
	return models.CurrentUser{ID: userID, RoleID: claims.RoleID, SessionID: claims.ID}, nil
}

func (m *JWTManager) ParseRefresh(token string) (int64, string, error) {
	claims, userID, err := m.parse(token, refreshTokenType)
	if err != nil {
		return 0, "", err
	}
	return userID, claims.ID, nil
}

func (m *JWTManager) registeredClaims(userID int64, sessionID string, ttl time.Duration) jwt.RegisteredClaims {
	now := time.Now()
	return jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ID:        sessionID,
		Issuer:    m.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}
}

func (m *JWTManager) sign(claims tokenClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("sign %s token: %w", claims.Type, err)
	}
	return token, nil
}

// parse проверяет подпись, срок, issuer и тип токена и возвращает id пользователя из sub.
func (m *JWTManager) parse(token, wantType string) (*tokenClaims, int64, error) {
	var claims tokenClaims
	_, err := jwt.ParseWithClaims(token, &claims,
		func(*jwt.Token) (any, error) { return m.secret, nil },
		// Принимаем только наш алгоритм: иначе можно подсунуть токен, подписанный по-другому.
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(m.issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %w", models.ErrInvalidToken, err)
	}
	if claims.Type != wantType {
		return nil, 0, fmt.Errorf("%w: expected %s token", models.ErrInvalidToken, wantType)
	}
	if claims.ID == "" {
		return nil, 0, fmt.Errorf("%w: missing session id", models.ErrInvalidToken)
	}

	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: invalid subject", models.ErrInvalidToken)
	}
	return &claims, userID, nil
}

type tokenClaims struct {
	jwt.RegisteredClaims
	Type   string `json:"type"`
	RoleID int64  `json:"role_id,omitempty"` // только в access
}
