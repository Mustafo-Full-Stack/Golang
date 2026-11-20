package habit

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
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

func (r *Repository) ListByUser(ctx context.Context, userID int64, limit, offset int) ([]models.Habit, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, name, created_at FROM habits
		WHERE user_id = $1 ORDER BY id
		LIMIT $2 OFFSET $3`,
		userID, limit, offset,
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

func (r *Repository) CountByUser(ctx context.Context, userID int64) (int64, error) {
	return r.countByID(ctx, `SELECT count(*) FROM habits WHERE user_id = $1`, userID)
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.Habit, error) {
	var h models.Habit

	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, name, created_at FROM habits WHERE id = $1`,
		id,
	).Scan(&h.ID, &h.UserID, &h.Name, &h.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *Repository) CreateCheck(ctx context.Context, habitID int64, date time.Time) error {
	tag, err := r.db.Exec(ctx,
		`INSERT INTO habit_checks (habit_id, date) VALUES ($1, $2)
		ON CONFLICT (habit_id, date) DO NOTHING`,
		habitID, date,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return models.ErrAlreadyChecked
	}

	return nil
}

func (r *Repository) ListChecks(ctx context.Context, habitID int64, limit, offset int) ([]models.HabitCheck, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, habit_id, date FROM habit_checks
		WHERE habit_id = $1 ORDER BY date
		LIMIT $2 OFFSET $3`,
		habitID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.HabitCheck

	for rows.Next() {
		var check models.HabitCheck

		if err := rows.Scan(&check.ID, &check.HabitID, &check.Date); err != nil {
			return nil, err
		}

		list = append(list, check)
	}

	return list, rows.Err()
}

func (r *Repository) CountChecks(ctx context.Context, habitID int64) (int64, error) {
	return r.countByID(ctx, `SELECT count(*) FROM habit_checks WHERE habit_id = $1`, habitID)
}

func (r *Repository) countByID(ctx context.Context, query string, id int64) (int64, error) {
	var total int64

	if err := r.db.QueryRow(ctx, query, id).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
