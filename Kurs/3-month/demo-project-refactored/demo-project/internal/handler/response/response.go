package response

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"demo-project/internal/models"
)

type errorBody struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("encode json response", "error", err)
	}
}

// Error отвечает клиенту текстом message, а причину cause пишет только в лог.
func Error(w http.ResponseWriter, r *http.Request, status int, message string, cause error) {
	level := slog.LevelWarn
	if status >= http.StatusInternalServerError {
		level = slog.LevelError
	}
	slog.Log(r.Context(), level, "request failed",
		"method", r.Method,
		"path", r.URL.Path,
		"status", status,
		"error", cause,
	)
	JSON(w, status, errorBody{Error: message})
}

func ServiceError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidInput):
		// Текст ошибки собран в models для клиента, например "invalid input: invalid email".
		Error(w, r, http.StatusBadRequest, err.Error(), err)
	case errors.Is(err, models.ErrInvalidCredentials):
		Error(w, r, http.StatusUnauthorized, models.ErrInvalidCredentials.Error(), err)
	case errors.Is(err, models.ErrInvalidToken):
		Error(w, r, http.StatusUnauthorized, models.ErrInvalidToken.Error(), err)
	case errors.Is(err, models.ErrEmailTaken):
		Error(w, r, http.StatusConflict, models.ErrEmailTaken.Error(), err)
	case errors.Is(err, models.ErrNotFound):
		Error(w, r, http.StatusNotFound, models.ErrNotFound.Error(), err)
	case errors.Is(err, models.ErrForbidden):
		Error(w, r, http.StatusForbidden, models.ErrForbidden.Error(), err)
	case errors.Is(err, models.ErrFileTooLarge):
		Error(w, r, http.StatusRequestEntityTooLarge, models.ErrFileTooLarge.Error(), err)
	case errors.Is(err, models.ErrUnsupportedFileType):
		Error(w, r, http.StatusUnsupportedMediaType, models.ErrUnsupportedFileType.Error(), err)
	case errors.Is(err, models.ErrFileInUse):
		Error(w, r, http.StatusConflict, models.ErrFileInUse.Error(), err)
	case errors.Is(err, models.ErrFileNotAvailable):
		Error(w, r, http.StatusBadRequest, models.ErrFileNotAvailable.Error(), err)
	default:
		Error(w, r, http.StatusInternalServerError, "internal server error", err)
	}
}
