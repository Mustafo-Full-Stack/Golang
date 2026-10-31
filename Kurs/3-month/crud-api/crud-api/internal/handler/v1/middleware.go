package v1

import (
	"crud-api/internal/handler/response"
	"crud-api/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	bearerPrefix = "Bearer "
	tokenKey     = "token"
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

	c.Set(tokenKey, token)
	c.Set(userIDKey, userID)
	c.Next()
}

func currentToken(c *gin.Context) string {
	return c.GetString(tokenKey)
}

func currentUserID(c *gin.Context) models.ID {
	return c.MustGet(userIDKey).(models.ID)
}
