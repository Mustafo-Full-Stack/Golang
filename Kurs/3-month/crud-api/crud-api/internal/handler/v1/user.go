package v1

import (
	"crud-api/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	ID        models.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) getMe(c *gin.Context) {
	user, err := h.srvs.UserService.GetByID(c.Request.Context(), currentUserID(c))
	if err != nil {
		writeError(c, "get me", err)
		return
	}

	c.JSON(http.StatusOK, userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}
