package v1

import (
	"crud-api/internal/handler/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	bearerPrefix = "Bearer "
	userIDKey    = "userID"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	token := strings.TrimSpace(strings.TrimPrefix(header, bearerPrefix))
	if !strings.HasPrefix(header, bearerPrefix) || token == "" {
		response.Error(c, http.StatusUnauthorized, "missing bearer token")
		return
	}

	userID, err := h.srvs.AuthService.Authenticate(c.Request.Context(), token)
	if err != nil {
		writeError(c, "auth middleware", err)
		return
	}

	c.Set(userIDKey, userID)
	c.Next()
}

func currentUserID(c *gin.Context) int64 {
	return c.GetInt64(userIDKey)
}
