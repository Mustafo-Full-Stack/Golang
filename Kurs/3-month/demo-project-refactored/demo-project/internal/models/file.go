package models

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const maxFileNameRunes = 255

// File — метаданные объекта в S3. PostID == nil — файл загружен, но не прикреплён к посту.
// URL в БД не хранится: сервис подписывает его при каждом чтении.
type File struct {
	ID          int64
	UserID      int64
	PostID      *int64
	ObjectKey   string
	Name        string
	ContentType string
	SizeBytes   int64
	URL         string
}

type UploadFileInput struct {
	Name    string
	Content io.Reader
}

func (in *UploadFileInput) Validate() error {
	// filepath.Base отрезает путь: "../../etc/passwd" → "passwd".
	in.Name = filepath.Base(strings.TrimSpace(in.Name))

	if in.Name == "." || in.Name == "/" {
		return fmt.Errorf("%w: file name is required", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Name) > maxFileNameRunes {
		return fmt.Errorf("%w: file name must be at most %d characters", ErrInvalidInput, maxFileNameRunes)
	}
	if in.Content == nil {
		return fmt.Errorf("%w: file is required", ErrInvalidInput)
	}
	return nil
}
