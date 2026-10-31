package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

func (s *Service) Create(ctx context.Context, user models.CurrentUser, in models.CreatePostInput) (*models.Post, error) {
	if !user.CanCreate() {
		return nil, models.ErrForbidden
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	post := &models.Post{
		UserID:  user.ID,
		Title:   in.Title,
		Content: in.Content,
	}
	if err := s.posts.Create(ctx, post, in.FileIDs); err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}

	// Перечитываем пост, чтобы вернуть его вместе с файлами и ссылками на них.
	return s.GetByID(ctx, post.ID)
}
