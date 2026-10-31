package service

import (
	"context"

	"demo-project/internal/models"
)

// Сигнатуры: ctx, затем user (если важно, кто действует), затем данные запроса —
// 1–2 простых значения отдельными параметрами, больше — одной структурой XxxInput.
type Services struct {
	Auth AuthService
	User UserService
	Post PostService
	File FileService
}

type AuthService interface {
	Register(ctx context.Context, in models.RegisterInput) (*models.User, error)
	Login(ctx context.Context, in models.LoginInput) (*models.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*models.TokenPair, error)
	Logout(ctx context.Context, user models.CurrentUser) error
	Authenticate(ctx context.Context, accessToken string) (models.CurrentUser, error)

	ListSessions(ctx context.Context, user models.CurrentUser) ([]models.Session, error)
	RevokeSession(ctx context.Context, user models.CurrentUser, sessionID string) error
	RevokeOtherSessions(ctx context.Context, user models.CurrentUser) error
}

type UserService interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
}

type PostService interface {
	Create(ctx context.Context, user models.CurrentUser, in models.CreatePostInput) (*models.Post, error)
	List(ctx context.Context, limit, offset int) ([]models.Post, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	Update(ctx context.Context, user models.CurrentUser, in models.UpdatePostInput) error
	Delete(ctx context.Context, user models.CurrentUser, id int64) error
}

type FileService interface {
	Upload(ctx context.Context, user models.CurrentUser, in models.UploadFileInput) (*models.File, error)
	Get(ctx context.Context, user models.CurrentUser, id int64) (*models.File, error)
	Delete(ctx context.Context, user models.CurrentUser, id int64) error
}
