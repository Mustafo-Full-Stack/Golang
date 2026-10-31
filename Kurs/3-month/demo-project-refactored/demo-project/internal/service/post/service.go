package post

import (
	"context"
	"fmt"

	"demo-project/internal/models"
	"demo-project/internal/repository"
	"demo-project/internal/service"
)

var _ service.PostService = (*Service)(nil)

type Service struct {
	posts   repository.PostRepository
	files   repository.FileRepository
	storage repository.ObjectStorage
}

func NewService(posts repository.PostRepository, files repository.FileRepository, storage repository.ObjectStorage) *Service {
	return &Service{posts: posts, files: files, storage: storage}
}

// getPostForEdit возвращает пост, если пользователь может его менять:
// ErrNotFound — поста нет, ErrForbidden — не автор, не модератор и не админ.
func (s *Service) getPostForEdit(ctx context.Context, user models.CurrentUser, id int64) (*models.Post, error) {
	post, err := s.posts.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !user.CanEditPost(post.UserID) {
		return nil, models.ErrForbidden
	}
	return post, nil
}

// attachFiles загружает файлы всех постов одним запросом и подписывает ссылки на них.
func (s *Service) attachFiles(ctx context.Context, posts []models.Post) error {
	if len(posts) == 0 {
		return nil
	}

	ids := make([]int64, 0, len(posts))
	for _, post := range posts {
		ids = append(ids, post.ID)
	}

	files, err := s.files.ListByPostIDs(ctx, ids)
	if err != nil {
		return fmt.Errorf("list post files: %w", err)
	}

	filesByPost := make(map[int64][]models.File)
	for _, file := range files {
		file.URL, err = s.storage.PresignedGetURL(ctx, file.ObjectKey, file.Name)
		if err != nil {
			return fmt.Errorf("sign file url: %w", err)
		}
		filesByPost[*file.PostID] = append(filesByPost[*file.PostID], file)
	}

	for i := range posts {
		posts[i].Files = filesByPost[posts[i].ID]
	}
	return nil
}
