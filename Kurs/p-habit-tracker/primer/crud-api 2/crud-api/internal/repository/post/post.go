package post

import (
	"context"
	"crud-api/internal/models"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

const uniqueViolationCode = "23505"

func (p *postRepository) CreatePost(ctx context.Context, post models.Post) (int64, error) {
	var id int64
	err := p.pool.QueryRow(ctx,
		`INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`,
		post.UserID, post.Title, post.Content,
	).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, models.ErrPostAlreadyExists
		}
		return 0, fmt.Errorf("create post: %w", err)
	}

	return id, nil
}
