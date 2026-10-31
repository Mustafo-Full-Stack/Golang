package post

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"demo-project/internal/models"
)

const selectColumns = `id, user_id, title, content, created_at, updated_at`

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// Create вставляет пост и прикрепляет к нему файлы в одной транзакции.
// Если хотя бы один файл чужой, не существует или уже прикреплён — откатывается всё.
func (r *Repository) Create(ctx context.Context, p *models.Post, fileIDs []int64) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		// После успешного Commit откатывать уже нечего: Rollback вернёт ErrTxClosed.
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			err = errors.Join(err, fmt.Errorf("rollback tx: %w", rbErr))
		}
	}()

	err = tx.QueryRow(ctx,
		`INSERT INTO posts (user_id, title, content)
		 VALUES ($1, $2, $3)
		 RETURNING id, created_at, updated_at`,
		p.UserID, p.Title, p.Content,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("insert post: %w", err)
	}

	if len(fileIDs) > 0 {
		tag, err := tx.Exec(ctx,
			`UPDATE files
			    SET post_id = $1
			  WHERE id = ANY($2)
			    AND user_id = $3
			    AND post_id IS NULL`,
			p.ID, fileIDs, p.UserID,
		)
		if err != nil {
			return fmt.Errorf("attach files: %w", err)
		}
		if tag.RowsAffected() != int64(len(fileIDs)) {
			return models.ErrFileNotAvailable
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.Post, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+selectColumns+`
		 FROM posts
		 ORDER BY created_at DESC, id DESC
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select posts: %w", err)
	}
	return collectPosts(rows)
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	var total int64
	if err := r.db.QueryRow(ctx, `SELECT count(*) FROM posts`).Scan(&total); err != nil {
		return 0, fmt.Errorf("count posts: %w", err)
	}
	return total, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	rows, err := r.db.Query(ctx, `SELECT `+selectColumns+` FROM posts WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("select post: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan post: %w", err)
	}
	post := m.toModel()
	return &post, nil
}

func (r *Repository) Update(ctx context.Context, p *models.Post) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE posts
		 SET title = $2, content = $3, updated_at = now()
		 WHERE id = $1`,
		p.ID, p.Title, p.Content,
	)
	if err != nil {
		return fmt.Errorf("update post: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete post: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func collectPosts(rows pgx.Rows) ([]models.Post, error) {
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan posts: %w", err)
	}

	posts := make([]models.Post, 0, len(items))
	for _, m := range items {
		posts = append(posts, m.toModel())
	}
	return posts, nil
}
