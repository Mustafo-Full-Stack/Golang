package post

import (
	"crud-api/internal/repository"
	"crud-api/internal/service"
)

var _ service.PostService = (*postService)(nil)

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) service.PostService {
	return &postService{postRepo: postRepo}
}
