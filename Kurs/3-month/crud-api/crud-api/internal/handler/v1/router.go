package v1

import (
	"crud-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srvs *service.Services
}

func NewHandler(srvs *service.Services) *Handler {
	return &Handler{
		srvs: srvs,
	}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.POST("/logout", h.authMiddleware, h.logout)
	}

	users := rg.Group("/users", h.authMiddleware)
	{
		users.GET("/me", h.getMe)
	}
}
