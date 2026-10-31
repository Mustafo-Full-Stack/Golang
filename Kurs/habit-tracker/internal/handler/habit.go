package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"habit-tracker/internal/models"
)

type createHabitRequest struct {
	Name string `json:"name"`
}

type habitResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type checkResponse struct {
	Date string `json:"date"`
}

// createHabit создаёт привычку текущего пользователя.
func (h *Handler) createHabit(w http.ResponseWriter, r *http.Request) {
	var req createHabitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	habit, err := h.habit.Create(r.Context(), models.CreateHabitInput{
		UserID: currentUserID(r),
		Name:   req.Name,
	})
	if err != nil {
		writeErr(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, habitResponse{
		ID:        habit.ID,
		Name:      habit.Name,
		CreatedAt: habit.CreatedAt,
	})
}

// listHabits возвращает привычки текущего пользователя.
func (h *Handler) listHabits(w http.ResponseWriter, r *http.Request) {
	habits, err := h.habit.List(r.Context(), currentUserID(r))
	if err != nil {
		writeErr(w, err)
		return
	}

	resp := make([]habitResponse, 0, len(habits))
	for _, habit := range habits {
		resp = append(resp, habitResponse{
			ID:        habit.ID,
			Name:      habit.Name,
			CreatedAt: habit.CreatedAt,
		})
	}

	writeJSON(w, http.StatusOK, resp)
}

// checkHabit отмечает привычку на сегодня.
func (h *Handler) checkHabit(w http.ResponseWriter, r *http.Request) {
	habitID, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid habit id")
		return
	}

	if err := h.habit.Check(r.Context(), currentUserID(r), habitID); err != nil {
		writeErr(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"status": "checked"})
}

// listChecks возвращает отметки привычки по дням.
func (h *Handler) listChecks(w http.ResponseWriter, r *http.Request) {
	habitID, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid habit id")
		return
	}

	checks, err := h.habit.Checks(r.Context(), currentUserID(r), habitID)
	if err != nil {
		writeErr(w, err)
		return
	}

	resp := make([]checkResponse, 0, len(checks))
	for _, check := range checks {
		resp = append(resp, checkResponse{Date: check.Date.Format("2006-01-02")})
	}

	writeJSON(w, http.StatusOK, resp)
}

// parseID читает {id} из адреса.
func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}
