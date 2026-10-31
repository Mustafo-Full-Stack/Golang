package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"habit-tracker/internal/models"
	"habit-tracker/internal/service/auth"
	"habit-tracker/internal/service/habit"
)

// Handler принимает HTTP-запросы и передаёт их в service.
type Handler struct {
	auth  *auth.Service
	habit *habit.Service
}

func New(authService *auth.Service, habitService *habit.Service) *Handler {
	return &Handler{auth: authService, habit: habitService}
}

// Routes описывает маршруты приложения.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sign-up", h.signUp)
	mux.HandleFunc("POST /login", h.login)

	mux.HandleFunc("POST /habits", h.authMiddleware(h.createHabit))
	mux.HandleFunc("GET /habits", h.authMiddleware(h.listHabits))
	mux.HandleFunc("POST /habits/{id}/check", h.authMiddleware(h.checkHabit))
	mux.HandleFunc("GET /habits/{id}/checks", h.authMiddleware(h.listChecks))

	return mux
}

// writeErr переводит ошибку service в HTTP-код.
func writeErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidInput),
		errors.Is(err, models.ErrPasswordTooShort),
		errors.Is(err, models.ErrPasswordTooLong):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, models.ErrInvalidToken), errors.Is(err, models.ErrInvalidCredentials):
		writeError(w, http.StatusUnauthorized, err.Error())
	case errors.Is(err, models.ErrForbidden):
		writeError(w, http.StatusForbidden, err.Error())
	case errors.Is(err, models.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, models.ErrUserExists), errors.Is(err, models.ErrAlreadyChecked):
		writeError(w, http.StatusConflict, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
