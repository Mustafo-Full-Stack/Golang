package v1

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
)

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	current := middleware.CurrentUser(r.Context())

	user, err := h.services.User.GetByID(r.Context(), current.ID)
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newUserResponse(user))
}
