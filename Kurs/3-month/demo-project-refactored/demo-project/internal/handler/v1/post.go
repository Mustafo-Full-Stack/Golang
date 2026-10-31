package v1

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
)

const defaultListLimit = 20

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req createPostRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	user := middleware.CurrentUser(r.Context())
	post, err := h.services.Post.Create(r.Context(), user, req.toInput())
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusCreated, newPostResponse(post))
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	limit, ok := queryInt(w, r, "limit", defaultListLimit)
	if !ok {
		return
	}
	offset, ok := queryInt(w, r, "offset", 0)
	if !ok {
		return
	}

	posts, total, err := h.services.Post.List(r.Context(), limit, offset)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newPostListResponse(posts, limit, offset, total))
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	post, err := h.services.Post.GetByID(r.Context(), id)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newPostResponse(post))
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var req updatePostRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	user := middleware.CurrentUser(r.Context())
	if err := h.services.Post.Update(r.Context(), user, req.toInput(id)); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	if err := h.services.Post.Delete(r.Context(), user, id); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
