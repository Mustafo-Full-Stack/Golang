package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

// Delete удаляет пост вместе с его файлами: записи о файлах уходят каскадом в БД,
// а сами объекты удаляем из S3 после этого.
func (s *Service) Delete(ctx context.Context, user models.CurrentUser, id int64) error {
	if _, err := s.getPostForEdit(ctx, user, id); err != nil {
		return fmt.Errorf("get post: %w", err)
	}

	files, err := s.files.ListByPostIDs(ctx, []int64{id})
	if err != nil {
		return fmt.Errorf("list post files: %w", err)
	}

	if err := s.posts.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete post: %w", err)
	}

	for _, file := range files {
		if err := s.storage.Delete(ctx, file.ObjectKey); err != nil {
			return fmt.Errorf("delete file object: %w", err)
		}
	}
	return nil
}
