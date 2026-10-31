package v1

import (
	"errors"
	"log/slog"
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
	"demo-project/internal/models"
)

// UploadFile принимает multipart/form-data с полем file.
func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Не даём прислать тело больше лимита; +1 МБ — запас на служебные части multipart.
	r.Body = http.MaxBytesReader(w, r.Body, h.cfg.MaxUploadBytes+1<<20)

	content, header, err := r.FormFile("file")
	if err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			response.ServiceError(w, r, models.ErrFileTooLarge)
			return
		}
		response.Error(w, r, http.StatusBadRequest, "file field is required", err)
		return
	}
	defer func() {
		if err := content.Close(); err != nil {
			slog.WarnContext(r.Context(), "close uploaded file", "error", err)
		}
	}()

	user := middleware.CurrentUser(r.Context())
	file, err := h.services.File.Upload(r.Context(), user, models.UploadFileInput{
		Name:    header.Filename,
		Content: content,
	})
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusCreated, newFileResponse(file))
}

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	file, err := h.services.File.Get(r.Context(), user, id)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newFileResponse(file))
}

func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	if err := h.services.File.Delete(r.Context(), user, id); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
