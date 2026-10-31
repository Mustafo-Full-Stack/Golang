package repository

import (
	"context"
	"io"
	"time"

	"demo-project/internal/models"
)

type Repositories struct {
	User    UserRepository
	Session SessionRepository // Postgres: список сессий пользователя
	Token   TokenRepository   // Redis: хеши текущих токенов сессии
	Post    PostRepository
	File    FileRepository
	Storage ObjectStorage // S3/MinIO: сами файлы
}

type UserRepository interface {
	Create(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, s *models.Session) error
	List(ctx context.Context, userID int64) ([]models.Session, error)
	Extend(ctx context.Context, id string, expiresAt time.Time) error
	Delete(ctx context.Context, userID int64, id string) error
	DeleteOthers(ctx context.Context, userID int64, keepID string) ([]string, error)
	DeleteExpired(ctx context.Context) (int64, error)
}

type TokenRepository interface {
	SavePair(ctx context.Context, sessionID string, pair *models.TokenPair, accessTTL, refreshTTL time.Duration) error
	IsAccessValid(ctx context.Context, sessionID, accessToken string) (bool, error)
	UseRefresh(ctx context.Context, sessionID, refreshToken string) (bool, error)
	Delete(ctx context.Context, sessionID string) error
	DeleteMany(ctx context.Context, sessionIDs []string) error
}

type PostRepository interface {
	Create(ctx context.Context, p *models.Post, fileIDs []int64) error
	List(ctx context.Context, limit, offset int) ([]models.Post, error)
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	Update(ctx context.Context, p *models.Post) error
	Delete(ctx context.Context, id int64) error
}

type FileRepository interface {
	Create(ctx context.Context, f *models.File) error
	GetByID(ctx context.Context, id int64) (*models.File, error)
	ListByPostIDs(ctx context.Context, postIDs []int64) ([]models.File, error)
	ListUnattachedBefore(ctx context.Context, before time.Time) ([]models.File, error)
	Delete(ctx context.Context, id int64) error
}

type ObjectStorage interface {
	Upload(ctx context.Context, key string, content io.Reader, size int64, contentType string) error
	PresignedGetURL(ctx context.Context, key, filename string) (string, error)
	Delete(ctx context.Context, key string) error
}
