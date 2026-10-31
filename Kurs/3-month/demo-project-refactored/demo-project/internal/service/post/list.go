package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

const maxListLimit = 100

func (s *Service) List(ctx context.Context, limit, offset int) ([]models.Post, int64, error) {
	if limit < 1 || limit > maxListLimit {
		return nil, 0, fmt.Errorf("%w: limit must be between 1 and %d", models.ErrInvalidInput, maxListLimit)
	}
	if offset < 0 {
		return nil, 0, fmt.Errorf("%w: offset must be zero or greater", models.ErrInvalidInput)
	}

	posts, err := s.posts.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list posts: %w", err)
	}

	total, err := s.posts.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("count posts: %w", err)
	}

	if err := s.attachFiles(ctx, posts); err != nil {
		return nil, 0, err
	}
	return posts, total, nil
}
