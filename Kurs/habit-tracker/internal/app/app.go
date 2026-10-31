package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"habit-tracker/internal/config"
	"habit-tracker/internal/handler"
	habitrepo "habit-tracker/internal/repository/habit"
	userrepo "habit-tracker/internal/repository/user"
	"habit-tracker/internal/service/auth"
	"habit-tracker/internal/service/habit"
	"habit-tracker/pkg/db/postgres"
)

// Run собирает все слои приложения и запускает HTTP-сервер.
func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	ctx := context.Background()
	pool, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}
	defer pool.Close()

	users := userrepo.NewRepository(pool)
	authService := auth.NewService(users, cfg.JWTSecret, cfg.JWTTTL)

	habits := habitrepo.NewRepository(pool)
	habitService := habit.NewService(habits)

	h := handler.New(authService, habitService)

	log.Printf("http server listening on :%s", cfg.HTTPPort)
	return http.ListenAndServe(":"+cfg.HTTPPort, h.Routes())
}
