package v1

import (
	"crud-api/internal/handler/response"
	"crud-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type registerResponse struct {
	ID models.ID `json:"id"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	userID, err := h.srvs.AuthService.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		writeError(c, "register", err)
		return
	}

	c.JSON(http.StatusCreated, registerResponse{ID: userID})
}

func (h *Handler) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.srvs.AuthService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		writeError(c, "login", err)
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

func (h *Handler) logout(c *gin.Context) {
	if err := h.srvs.AuthService.Logout(c.Request.Context(), currentToken(c)); err != nil {
		writeError(c, "logout", err)
		return
	}

	c.Status(http.StatusNoContent)
}
