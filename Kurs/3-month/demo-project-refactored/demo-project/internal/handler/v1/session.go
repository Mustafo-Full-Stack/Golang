package v1

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
)

func (h *Handler) ListSessions(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r.Context())

	sessions, err := h.services.Auth.ListSessions(r.Context(), user)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newSessionListResponse(sessions, user.SessionID))
}

func (h *Handler) RevokeSession(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r.Context())

	if err := h.services.Auth.RevokeSession(r.Context(), user, r.PathValue("id")); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) RevokeOtherSessions(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r.Context())

	if err := h.services.Auth.RevokeOtherSessions(r.Context(), user); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
