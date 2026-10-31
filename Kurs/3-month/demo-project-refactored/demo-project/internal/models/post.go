package models

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	maxTitleRunes   = 200
	maxContentRunes = 10_000
	maxPostFiles    = 10
)

type Post struct {
	ID        int64
	UserID    int64
	Title     string
	Content   string
	Files     []File
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePostInput struct {
	Title   string
	Content string
	FileIDs []int64
}

func (in *CreatePostInput) Validate() error {
	in.Title = strings.TrimSpace(in.Title)
	in.Content = strings.TrimSpace(in.Content)
	in.FileIDs = uniqueIDs(in.FileIDs) // один файл нельзя прикрепить дважды

	if err := validatePostText(in.Title, in.Content); err != nil {
		return err
	}
	if len(in.FileIDs) > maxPostFiles {
		return fmt.Errorf("%w: at most %d files per post", ErrInvalidInput, maxPostFiles)
	}
	for _, id := range in.FileIDs {
		if id <= 0 {
			return fmt.Errorf("%w: file id must be positive", ErrInvalidInput)
		}
	}
	return nil
}

type UpdatePostInput struct {
	ID      int64
	Title   string
	Content string
}

func (in *UpdatePostInput) Validate() error {
	in.Title = strings.TrimSpace(in.Title)
	in.Content = strings.TrimSpace(in.Content)

	if in.ID <= 0 {
		return fmt.Errorf("%w: invalid post id", ErrInvalidInput)
	}
	return validatePostText(in.Title, in.Content)
}

func validatePostText(title, content string) error {
	// Postgres не принимает NUL в text: без проверки такой запрос дал бы 500.
	if strings.ContainsRune(title, '\x00') {
		return fmt.Errorf("%w: title must not contain NUL", ErrInvalidInput)
	}
	if strings.ContainsRune(content, '\x00') {
		return fmt.Errorf("%w: content must not contain NUL", ErrInvalidInput)
	}
	if n := utf8.RuneCountInString(title); n == 0 || n > maxTitleRunes {
		return fmt.Errorf("%w: title must be 1 to %d characters", ErrInvalidInput, maxTitleRunes)
	}
	if n := utf8.RuneCountInString(content); n == 0 || n > maxContentRunes {
		return fmt.Errorf("%w: content must be 1 to %d characters", ErrInvalidInput, maxContentRunes)
	}
	return nil
}

func uniqueIDs(ids []int64) []int64 {
	var result []int64
	seen := make(map[int64]bool)
	for _, id := range ids {
		if !seen[id] {
			seen[id] = true
			result = append(result, id)
		}
	}
	return result
}
