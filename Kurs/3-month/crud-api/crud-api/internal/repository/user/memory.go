package user

import (
	"context"
	"crud-api/internal/models"
	"crud-api/internal/repository"
	"strings"
	"sync"
	"time"
)

var _ repository.UserRepository = (*userRepoInMemory)(nil)

type userRepoInMemory struct {
	mu     sync.RWMutex
	users  map[models.ID]models.User
	nextID models.ID
}

func NewUserRepoInMemory() *userRepoInMemory {
	return &userRepoInMemory{
		users:  make(map[models.ID]models.User),
		nextID: 1,
	}
}

func (r *userRepoInMemory) CreateUser(ctx context.Context, user *models.User) (models.ID, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existingUser := range r.users {
		if strings.EqualFold(existingUser.Email, user.Email) {
			return 0, models.ErrEmailAlreadyExists
		}
	}
	user.ID = r.nextID
	user.CreatedAt = time.Now()

	r.users[user.ID] = *user
	r.nextID++

	return user.ID, nil
}

func (r *userRepoInMemory) GetByID(ctx context.Context, id models.ID) (*models.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, models.ErrUserNotFound
	}

	return &user, nil
}

func (r *userRepoInMemory) GetAll(ctx context.Context) ([]models.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]models.User, 0, len(r.users))

	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepoInMemory) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if strings.EqualFold(user.Email, email) {
			return &user, nil
		}
	}

	return nil, models.ErrUserNotFound
}
