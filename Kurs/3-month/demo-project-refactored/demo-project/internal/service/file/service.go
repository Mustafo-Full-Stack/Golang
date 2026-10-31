package file

import (
	"context"
	"fmt"

	"demo-project/internal/models"
	"demo-project/internal/repository"
	"demo-project/internal/service"
)

var _ service.FileService = (*Service)(nil)

var allowedContentTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/webp":      true,
	"image/gif":       true,
	"application/pdf": true,
}

type Service struct {
	files          repository.FileRepository
	storage        repository.ObjectStorage
	maxUploadBytes int64
}

func NewService(files repository.FileRepository, storage repository.ObjectStorage, maxUploadBytes int64) *Service {
	return &Service{
		files:          files,
		storage:        storage,
		maxUploadBytes: maxUploadBytes,
	}
}

func (s *Service) signURL(ctx context.Context, file *models.File) error {
	url, err := s.storage.PresignedGetURL(ctx, file.ObjectKey, file.Name)
	if err != nil {
		return fmt.Errorf("sign file url: %w", err)
	}
	file.URL = url
	return nil
}
