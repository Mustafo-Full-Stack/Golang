package storage

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type Storage struct {
	client     *minio.Client
	bucket     string
	presignTTL time.Duration
}

func NewStorage(client *minio.Client, bucket string, presignTTL time.Duration) *Storage {
	return &Storage{
		client:     client,
		bucket:     bucket,
		presignTTL: presignTTL,
	}
}

func (s *Storage) Upload(ctx context.Context, key string, content io.Reader, size int64, contentType string) error {
	_, err := s.client.PutObject(ctx, s.bucket, key, content, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("upload object: %w", err)
	}
	return nil
}

func (s *Storage) PresignedGetURL(ctx context.Context, key, filename string) (string, error) {
	params := url.Values{}
	// Имя файла при скачивании; mime.FormatMediaType сам экранирует спецсимволы.
	if disposition := mime.FormatMediaType("inline", map[string]string{"filename": filename}); disposition != "" {
		params.Set("response-content-disposition", disposition)
	}

	u, err := s.client.PresignedGetObject(ctx, s.bucket, key, s.presignTTL, params)
	if err != nil {
		return "", fmt.Errorf("presign object: %w", err)
	}
	return u.String(), nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	if err := s.client.RemoveObject(ctx, s.bucket, key, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("delete object: %w", err)
	}
	return nil
}
