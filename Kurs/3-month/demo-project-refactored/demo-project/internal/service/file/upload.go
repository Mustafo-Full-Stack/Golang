package file

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"

	"demo-project/internal/models"
)

// Upload сохраняет файл в S3 и запись о нём в БД. Тип файла определяем по содержимому,
// а не по расширению: иначе под видом картинки можно загрузить что угодно.
func (s *Service) Upload(ctx context.Context, user models.CurrentUser, in models.UploadFileInput) (*models.File, error) {
	if !user.CanCreate() {
		return nil, models.ErrForbidden
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	data, err := io.ReadAll(in.Content)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: file is empty", models.ErrInvalidInput)
	}
	if int64(len(data)) > s.maxUploadBytes {
		return nil, models.ErrFileTooLarge
	}

	contentType := http.DetectContentType(data)
	if !allowedContentTypes[contentType] {
		return nil, fmt.Errorf("%w: %s", models.ErrUnsupportedFileType, contentType)
	}

	file := &models.File{
		UserID:      user.ID,
		ObjectKey:   fmt.Sprintf("users/%d/%s", user.ID, uuid.NewString()),
		Name:        in.Name,
		ContentType: contentType,
		SizeBytes:   int64(len(data)),
	}
	if err := s.storage.Upload(ctx, file.ObjectKey, bytes.NewReader(data), file.SizeBytes, contentType); err != nil {
		return nil, fmt.Errorf("upload file: %w", err)
	}

	if err := s.files.Create(ctx, file); err != nil {
		// Запись не сохранилась — удаляем и объект, чтобы он не остался в S3 без хозяина.
		if deleteErr := s.storage.Delete(ctx, file.ObjectKey); deleteErr != nil {
			return nil, errors.Join(fmt.Errorf("save file: %w", err), fmt.Errorf("delete object: %w", deleteErr))
		}
		return nil, fmt.Errorf("save file: %w", err)
	}

	if err := s.signURL(ctx, file); err != nil {
		return nil, err
	}
	return file, nil
}
