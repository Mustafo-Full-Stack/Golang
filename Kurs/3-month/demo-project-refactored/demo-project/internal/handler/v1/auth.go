package v1

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	"demo-project/internal/handler/response"
	"demo-project/internal/models"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	user, err := h.services.Auth.Register(r.Context(), req.toInput())
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusCreated, newUserResponse(user))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	tokens, err := h.services.Auth.Login(r.Context(), models.LoginInput{
		Email:      req.Email,
		Password:   req.Password,
		DeviceType: models.DeviceType(r.Header.Get("X-Device-Type")),
		IP:         clientIP(r),
		UserAgent:  r.UserAgent(),
	})
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newTokenPairResponse(tokens))
}

// Refresh принимает refresh-токен в заголовке Authorization: Bearer <refresh_token>.
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokens, err := h.services.Auth.Refresh(r.Context(), middleware.BearerToken(r))
	if err != nil {
		response.ServiceError(w, r, err)
		return
	}

	response.JSON(w, http.StatusOK, newTokenPairResponse(tokens))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r.Context())

	if err := h.services.Auth.Logout(r.Context(), user); err != nil {
		response.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
