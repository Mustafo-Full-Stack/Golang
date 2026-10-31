package file

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

// Delete удаляет неприкреплённый файл. Прикреплённый удаляется только вместе с постом.
func (s *Service) Delete(ctx context.Context, user models.CurrentUser, id int64) error {
	file, err := s.files.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get file: %w", err)
	}
	if file.PostID != nil {
		return models.ErrFileInUse
	}
	if !user.CanManageFile(file.UserID) {
		return models.ErrNotFound
	}

	if err := s.files.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete file: %w", err)
	}
	if err := s.storage.Delete(ctx, file.ObjectKey); err != nil {
		return fmt.Errorf("delete file object: %w", err)
	}
	return nil
}
