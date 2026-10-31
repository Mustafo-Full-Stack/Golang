package file

import (
	"context"
	"fmt"

	"demo-project/internal/models"
)

// Get отдаёт прикреплённый к посту файл любому пользователю,
// а неприкреплённый — только владельцу или админу. Остальным — ErrNotFound,
// чтобы не раскрывать, что такой файл вообще есть.
func (s *Service) Get(ctx context.Context, user models.CurrentUser, id int64) (*models.File, error) {
	file, err := s.files.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}
	if file.PostID == nil && !user.CanManageFile(file.UserID) {
		return nil, models.ErrNotFound
	}

	if err := s.signURL(ctx, file); err != nil {
		return nil, err
	}
	return file, nil
}
