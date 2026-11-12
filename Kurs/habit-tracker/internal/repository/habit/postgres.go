package habit

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"habit-tracker/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, h *models.Habit) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO habits (user_id, name) VALUES ($1, $2)
		RETURNING id, created_at`,
		h.UserID, h.Name,
	).Scan(&h.ID, &h.CreatedAt)
}

func (r *Repository) ListByUser(ctx context.Context, userID int64) ([]models.Habit, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, name, created_at FROM habits
			 WHERE user_id = 1$ ORDER BY id`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Habit

	for rows.Next() {
		var h models.Habit

		if err := rows.Scan(&h.ID, &h.UserID, &h.Name, &h.CreatedAt); err != nil {
			return nil, err
		}

		list = append(list, h)
	}

	return list, rows.Err()
}
