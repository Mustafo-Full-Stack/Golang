package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	post, err := s.posts.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get post: %w", err)
	}

	posts := []models.Post{*post}
	if err := s.attachFiles(ctx, posts); err != nil {
		return nil, err
	}
	return &posts[0], nil
}
