package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

func (s *Service) Update(ctx context.Context, user models.CurrentUser, in models.UpdatePostInput) error {
	if err := in.Validate(); err != nil {
		return err
	}

	post, err := s.getPostForEdit(ctx, user, in.ID)
	if err != nil {
		return fmt.Errorf("get post: %w", err)
	}

	post.Title = in.Title
	post.Content = in.Content
	if err := s.posts.Update(ctx, post); err != nil {
		return fmt.Errorf("update post: %w", err)
	}
	return nil
}
