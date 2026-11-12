package v1

import (
	"crud-api/internal/handler/response"
	"crud-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createPostResponse struct {
	ID int64 `json:"id"`
}

func (h *Handler) createPost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	postID, err := h.srvs.PostService.CreatePost(c.Request.Context(), models.Post{
		UserID:  currentUserID(c),
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		writeError(c, "create post", err)
		return
	}

	c.JSON(http.StatusCreated, createPostResponse{ID: postID})
}
