package app

import (
	"context"
	"crud-api/internal/config"
	"crud-api/internal/handler"
	postrepo "crud-api/internal/repository/post"
	userrepo "crud-api/internal/repository/user"
	"crud-api/internal/service"
	"crud-api/internal/service/auth"
	"crud-api/internal/service/post"
	jwttoken "crud-api/internal/service/token/jwt"
	"crud-api/internal/service/user"
	"crud-api/pkg/db/postgres"
	"crud-api/pkg/logger"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	slog.SetDefault(logger.New(cfg.Logger))

	if cfg.ServerConfigs.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	pool, err := postgres.New(context.Background(), cfg.Postgres)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}
	defer pool.Close()

	userRepo := userrepo.NewUserRepo(pool)
	postRepo := postrepo.NewPostRepository(pool)

	tokens := jwttoken.NewTokenManager(cfg.JWT.Secret, cfg.JWT.AccessTTL, cfg.JWT.RefreshTTL)

	srvs := &service.Services{
		AuthService: auth.NewAuthService(userRepo, tokens),
		UserService: user.NewUserService(userRepo),
		PostService: post.NewPostService(postRepo),
	}
	h := handler.NewHandler(srvs)

	server := &http.Server{
		Addr:              ":" + cfg.ServerConfigs.Port,
		Handler:           h.InitRoutes(),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	// Запускаем сервер в отдельной горутине, чтобы main мог дальше ждать сигнал остановки.
	go func() {
		log.Printf("http server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	}()

	// Ждём Ctrl+C (SIGINT) или SIGTERM от Docker/Kubernetes.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down http server")

	// Даём текущим запросам до shutdownTimeout, чтобы завершиться.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}

	log.Println("server stopped")
	return nil
}
