package post

import (
	"context"
	"crud-api/internal/models"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	maxTitleLength   = 70
	maxContentLength = 500
)

func (p *postService) CreatePost(ctx context.Context, post models.Post) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	if post.Title == "" || utf8.RuneCountInString(post.Title) > maxTitleLength {
		return 0, models.ErrInvalidPostTitle
	}
	if post.Content == "" || utf8.RuneCountInString(post.Content) > maxContentLength {
		return 0, models.ErrInvalidPostContent
	}

	id, err := p.postRepo.CreatePost(ctx, post)
	if err != nil {
		if errors.Is(err, models.ErrPostAlreadyExists) {
			return 0, err
		}
		return 0, fmt.Errorf("failed to create post: %w", err)
	}

	return id, nil
}
