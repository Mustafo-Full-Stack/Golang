package handler

import (
	v1 "crud-api/internal/handler/v1"
	"crud-api/internal/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	srvs *service.Services
}

func NewHandler(srvs *service.Services) *handler {
	return &handler{
		srvs: srvs,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")
	v1.NewHandler(h.srvs).InitRoutes(api.Group("/v1"))

	return router
}
