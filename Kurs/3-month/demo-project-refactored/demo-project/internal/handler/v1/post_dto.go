package v1

import (
	"time"

	"demo-project/internal/models"
)

type createPostRequest struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	FileIDs []int64 `json:"file_ids"`
}

func (req createPostRequest) toInput() models.CreatePostInput {
	return models.CreatePostInput{
		Title:   req.Title,
		Content: req.Content,
		FileIDs: req.FileIDs,
	}
}

type updatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (req updatePostRequest) toInput(id int64) models.UpdatePostInput {
	return models.UpdatePostInput{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}
}

type postResponse struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Files     []fileResponse `json:"files"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func newPostResponse(p *models.Post) postResponse {
	return postResponse{
		ID:        p.ID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Files:     newFilesResponse(p.Files),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// newPostsResponse всегда возвращает массив: пустая лента — [], а не null.
func newPostsResponse(posts []models.Post) []postResponse {
	resp := make([]postResponse, 0, len(posts))
	for i := range posts {
		resp = append(resp, newPostResponse(&posts[i]))
	}
	return resp
}

type postListResponse struct {
	Items  []postResponse `json:"items"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Total  int64          `json:"total"`
}

func newPostListResponse(posts []models.Post, limit, offset int, total int64) postListResponse {
	return postListResponse{
		Items:  newPostsResponse(posts),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}
