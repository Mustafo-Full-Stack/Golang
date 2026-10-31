package v1

import "demo-project/internal/service"

type Config struct {
	MaxUploadBytes int64
}

// Каждая ручка делает три шага: разобрать запрос → вызвать сервис → ответить.
type Handler struct {
	services *service.Services
	cfg      Config
}

func NewHandler(services *service.Services, cfg Config) *Handler {
	return &Handler{services: services, cfg: cfg}
}
