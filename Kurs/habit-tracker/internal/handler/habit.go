package handler

import (
	"encoding/json"
	"habit-tracker/internal/models"
	"net/http"
	"strconv"
	"time"
)

const defaultListLimit = 20

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

type listResponse[T any] struct {
	Items  []T   `json:"items"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

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

func (h *Handler) listHabits(w http.ResponseWriter, r *http.Request) {
	limit, offset, ok := parsePagination(w, r)
	if !ok {
		return
	}

	habits, total, err := h.habit.List(r.Context(), currentUserID(r), limit, offset)
	if err != nil {
		writeErr(w, err)
		return
	}

	items := make([]habitResponse, 0, len(habits))

	for _, habit := range habits {
		items = append(items, habitResponse{
			ID:        habit.ID,
			Name:      habit.Name,
			CreatedAt: habit.CreatedAt,
		})
	}

	writeJSON(w, http.StatusOK, listResponse[habitResponse]{
		Items:  items,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	})
}

func (h *Handler) checkHabit(w http.ResponseWriter, r *http.Request) {
	habitID, err := parseID(r)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid habit ID")
		return
	}

	if err := h.habit.Check(r.Context(), currentUserID(r), habitID); err != nil {
		writeErr(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"status": "checked"})
}

func (h *Handler) listChecks(w http.ResponseWriter, r *http.Request) {
	habitID, err := parseID(r)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid habit ID")
		return
	}

	limit, offset, ok := parsePagination(w, r)
	if !ok {
		return
	}

	checks, total, err := h.habit.Checks(r.Context(), currentUserID(r), habitID, limit, offset)

	if err != nil {
		writeErr(w, err)
		return
	}

	items := make([]checkResponse, 0, len(checks))

	for _, check := range checks {
		items = append(items, checkResponse{Date: check.Date.Format("2006-01-02")})
	}

	writeJSON(w, http.StatusOK, listResponse[checkResponse]{
		Items:  items,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	})
}

func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

func parsePagination(w http.ResponseWriter, r *http.Request) (int, int, bool) {
	limit, ok := queryInt(w, r, "limit", defaultListLimit)
	if !ok {
		return 0, 0, false
	}

	offset, ok := queryInt(w, r, "offset", 0)
	if !ok {
		return 0, 0, false
	}

	return limit, offset, true
}

func queryInt(w http.ResponseWriter, r *http.Request, name string, def int) (int, bool) {
	raw := r.URL.Query().Get(name)
	if raw == "" {
		return def, true
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid "+name)
		return 0, false
	}

	return val, true
}
