package post

import (
	"crud-api/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.PostRepository = (*postRepository)(nil)

type postRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *postRepository {
	return &postRepository{pool: pool}
}
