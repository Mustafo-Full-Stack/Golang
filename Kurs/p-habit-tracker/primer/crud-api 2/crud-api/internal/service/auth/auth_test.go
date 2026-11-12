package auth

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/repository"
	jwttoken "crud-api/internal/service/token/jwt"
	"errors"
	"strings"
	"testing"
	"time"
)

// fakeUserRepo — тестовая заглушка репозитория, чтобы проверять бизнес-логику без Postgres.
type fakeUserRepo struct {
	users map[string]models.User
}

var _ repository.UserRepository = (*fakeUserRepo)(nil)

func newFakeUserRepo() *fakeUserRepo {
	return &fakeUserRepo{users: make(map[string]models.User)}
}

func (r *fakeUserRepo) CreateUser(_ context.Context, user *models.User) (int64, error) {
	if _, ok := r.users[user.Email]; ok {
		return 0, models.ErrEmailAlreadyExists
	}
	user.ID = int64(len(r.users) + 1)
	r.users[user.Email] = *user
	return user.ID, nil
}

func (r *fakeUserRepo) GetByEmail(_ context.Context, email string) (*models.User, error) {
	user, ok := r.users[email]
	if !ok {
		return nil, models.ErrUserNotFound
	}
	return &user, nil
}

func (r *fakeUserRepo) GetByID(_ context.Context, id int64) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, models.ErrUserNotFound
}

func (r *fakeUserRepo) GetAll(_ context.Context) ([]models.User, error) {
	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		password string
		wantErr  error
	}{
		{
			name:     "success",
			userName: "Bob",
			email:    "bob@mail.com",
			password: "12345678",
		},
		{
			name:     "password counted in runes, not bytes",
			userName: "Bob",
			email:    "bob@mail.com",
			password: "пароль12",
		},
		{
			name:     "password too short",
			userName: "Bob",
			email:    "bob@mail.com",
			password: "1234567",
			wantErr:  models.ErrPasswordTooShort,
		},
		{
			name:     "multibyte password too short",
			userName: "Bob",
			email:    "bob@mail.com",
			password: "пароль1",
			wantErr:  models.ErrPasswordTooShort,
		},
		{
			name:     "password at max length",
			userName: "Bob",
			email:    "bob@mail.com",
			password: strings.Repeat("a", 72),
		},
		{
			name:     "password too long",
			userName: "Bob",
			email:    "bob@mail.com",
			password: strings.Repeat("a", 73),
			wantErr:  models.ErrPasswordTooLong,
		},
		{
			name:     "multibyte password too long in bytes",
			userName: "Bob",
			email:    "bob@mail.com",
			password: strings.Repeat("я", 37),
			wantErr:  models.ErrPasswordTooLong,
		},
		{
			name:     "blank name",
			userName: "   ",
			email:    "bob@mail.com",
			password: "12345678",
			wantErr:  models.ErrInvalidInput,
		},
		{
			name:     "empty password",
			userName: "Bob",
			email:    "bob@mail.com",
			wantErr:  models.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewAuthService(newFakeUserRepo(), jwttoken.NewTokenManager("test-secret", time.Hour, time.Hour))

			_, err := svc.Register(context.Background(), tt.userName, tt.email, tt.password)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	ctx := context.Background()
	svc := NewAuthService(newFakeUserRepo(), jwttoken.NewTokenManager("test-secret", time.Hour, time.Hour))

	userID, err := svc.Register(ctx, "Bob", "bob@mail.com", "12345678")
	if err != nil {
		t.Fatalf("register: %v", err)
	}

	tokens, err := svc.Login(ctx, "bob@mail.com", "12345678")
	if err != nil {
		t.Fatalf("login: %v", err)
	}

	tests := []struct {
		name    string
		token   string
		wantID  int64
		wantErr error
	}{
		{
			name:   "valid access token",
			token:  tokens.AccessToken,
			wantID: userID,
		},
		{
			name:    "refresh token is not accepted",
			token:   tokens.RefreshToken,
			wantErr: models.ErrInvalidToken,
		},
		{
			name:    "unknown token",
			token:   "unknown",
			wantErr: models.ErrInvalidToken,
		},
		{
			name:    "blank token",
			token:   "   ",
			wantErr: models.ErrInvalidToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := svc.Authenticate(ctx, tt.token)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if gotID != tt.wantID {
				t.Fatalf("id = %d, want %d", gotID, tt.wantID)
			}
		})
	}
}
