package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"demo-project/internal/models"
)

const uniqueViolation = "23505"

const selectUser = `
	SELECT u.id, u.email, u.password_hash, u.name, u.role_id, r.code AS role_code, u.created_at
	  FROM users u
	  JOIN roles r ON r.id = u.role_id`

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *models.User) error {
	err := r.db.QueryRow(ctx,
		`INSERT INTO users (email, password_hash, name, role_id)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, (SELECT code FROM roles WHERE id = $4)`,
		u.Email, u.PasswordHash, u.Name, u.RoleID,
	).Scan(&u.ID, &u.CreatedAt, &u.RoleCode)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
			return models.ErrEmailTaken
		}
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return r.get(ctx, selectUser+` WHERE u.id = $1`, id)
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.get(ctx, selectUser+` WHERE u.email = $1`, email)
}

func (r *Repository) get(ctx context.Context, query string, arg any) (*models.User, error) {
	rows, err := r.db.Query(ctx, query, arg)
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}
	user := m.toModel()
	return &user, nil
}
