package file

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"demo-project/internal/models"
)

const selectColumns = `id, user_id, post_id, object_key, original_name, content_type, size_bytes`

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, f *models.File) error {
	err := r.db.QueryRow(ctx,
		`INSERT INTO files (user_id, object_key, original_name, content_type, size_bytes)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		f.UserID, f.ObjectKey, f.Name, f.ContentType, f.SizeBytes,
	).Scan(&f.ID)
	if err != nil {
		return fmt.Errorf("insert file: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.File, error) {
	rows, err := r.db.Query(ctx, `SELECT `+selectColumns+` FROM files WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("select file: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan file: %w", err)
	}
	file := m.toModel()
	return &file, nil
}

func (r *Repository) ListByPostIDs(ctx context.Context, postIDs []int64) ([]models.File, error) {
	return r.list(ctx, `SELECT `+selectColumns+` FROM files WHERE post_id = ANY($1) ORDER BY id`, postIDs)
}

func (r *Repository) ListUnattachedBefore(ctx context.Context, before time.Time) ([]models.File, error) {
	return r.list(ctx, `SELECT `+selectColumns+` FROM files WHERE post_id IS NULL AND created_at < $1`, before)
}

// Delete удаляет только неприкреплённый файл: условие post_id IS NULL не даст
// удалить файл, который как раз в этот момент прикрепили к посту.
func (r *Repository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM files WHERE id = $1 AND post_id IS NULL`, id)
	if err != nil {
		return fmt.Errorf("delete file: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (r *Repository) list(ctx context.Context, query string, arg any) ([]models.File, error) {
	rows, err := r.db.Query(ctx, query, arg)
	if err != nil {
		return nil, fmt.Errorf("select files: %w", err)
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan files: %w", err)
	}

	files := make([]models.File, 0, len(items))
	for _, item := range items {
		files = append(files, item.toModel())
	}
	return files, nil
}
