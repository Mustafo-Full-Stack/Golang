package app

import (
	"context"
	"crud-api/internal/config"
	"crud-api/internal/handler"
	"crud-api/internal/repository"
	userrepo "crud-api/internal/repository/user"
	"crud-api/internal/service"
	"crud-api/internal/service/auth"
	"crud-api/internal/service/token"
	"crud-api/internal/service/token/memory"
	redistoken "crud-api/internal/service/token/redis"
	"crud-api/internal/service/user"
	"crud-api/pkg/db/postgres"
	"crud-api/pkg/db/redis"
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

func Run() (err error) {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	closeLogger, err := setupLogger(cfg)
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}
	defer closeLogger()

	if cfg.ServerConfigs.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	users, closeRepo, err := newUserRepo(context.Background(), cfg)
	if err != nil {
		return fmt.Errorf("init user repository: %w", err)
	}
	defer closeRepo()
	log.Printf("user repository: in-memory=%t", cfg.IsRepoInMemory)

	tokens, closeTokens, err := newTokenManager(context.Background(), cfg)
	if err != nil {
		return fmt.Errorf("init token manager: %w", err)
	}
	defer closeTokens()
	log.Printf("token manager: in-memory=%t", cfg.IsTokenInMemory)

	srvs := &service.Services{
		AuthService: auth.NewAuthService(users, tokens),
		UserService: user.NewUserService(users),
	}
	h := handler.NewHandler(srvs)

	server := &http.Server{
		Addr:              ":" + cfg.ServerConfigs.Port,
		Handler:           h.InitRoutes(),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("http server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err = <-errCh:
		return fmt.Errorf("http server: %w", err)
	case <-ctx.Done():
		log.Printf("shutting down http server")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}

	return nil
}

func newUserRepo(ctx context.Context, cfg *config.Config) (repository.UserRepository, func(), error) {
	if cfg.IsRepoInMemory {
		return userrepo.NewUserRepoInMemory(), func() {}, nil
	}

	pool, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, nil, fmt.Errorf("connect postgres: %w", err)
	}

	return userrepo.NewUserRepo(pool), pool.Close, nil
}

func newTokenManager(ctx context.Context, cfg *config.Config) (token.Manager, func(), error) {
	if cfg.IsTokenInMemory {
		return memory.NewTokenManager(cfg.TokenTTL), func() {}, nil
	}

	client, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		return nil, nil, fmt.Errorf("connect redis: %w", err)
	}

	closeClient := func() {
		if err := client.Close(); err != nil {
			log.Printf("close redis: %v", err)
		}
	}

	return redistoken.NewTokenManager(client, cfg.TokenTTL), closeClient, nil
}

func setupLogger(cfg *config.Config) (func(), error) {
	if !cfg.IsSlogEnabled {
		return func() {}, nil
	}

	l, closeFile, err := logger.New(cfg.Logger)
	if err != nil {
		return nil, err
	}
	slog.SetDefault(l)

	return func() {
		if err := closeFile(); err != nil {
			log.Printf("close log file: %v", err)
		}
	}, nil
}
