package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo-project/internal/config"
	"demo-project/internal/cron"
	"demo-project/internal/handler"
	v1 "demo-project/internal/handler/v1"
	"demo-project/internal/repository"
	filerepo "demo-project/internal/repository/file"
	postrepo "demo-project/internal/repository/post"
	sessionrepo "demo-project/internal/repository/session"
	"demo-project/internal/repository/storage"
	tokenrepo "demo-project/internal/repository/token"
	userrepo "demo-project/internal/repository/user"
	"demo-project/internal/service"
	authservice "demo-project/internal/service/auth"
	cleanupservice "demo-project/internal/service/cleanup"
	fileservice "demo-project/internal/service/file"
	postservice "demo-project/internal/service/post"
	userservice "demo-project/internal/service/user"
	"demo-project/pkg/logger"
	"demo-project/pkg/postgres"
	"demo-project/pkg/redis"
	"demo-project/pkg/s3"
)

func Run() (err error) {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log, closeLog, err := logger.New(logger.Config{
		Env:   cfg.App.Env,
		Level: cfg.App.LogLevel,
		File:  cfg.App.LogFile,
	})
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := closeLog(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close log file: %w", closeErr))
		}
	}()
	slog.SetDefault(log)

	// Пишем причину остановки, пока файл лога ещё открыт: defer выполняются в обратном порядке.
	defer func() {
		if err != nil {
			slog.Error("app stopped", "error", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	postgresPool, err := postgres.New(ctx, postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.Database,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		return err
	}
	defer postgresPool.Close()

	redisClient, err := redis.New(ctx, redis.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := redisClient.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close redis: %w", closeErr))
		}
	}()

	s3Client, err := s3.New(s3.Config{
		Endpoint:  cfg.S3.Endpoint,
		AccessKey: cfg.S3.AccessKey,
		SecretKey: cfg.S3.SecretKey,
		UseSSL:    cfg.S3.UseSSL,
	})
	if err != nil {
		return err
	}
	if err := s3.EnsureBucket(ctx, s3Client, cfg.S3.Bucket); err != nil {
		return err
	}

	jwtManager, err := authservice.NewJWTManager(authservice.JWTConfig{
		Secret:     []byte(cfg.JWT.Secret),
		Issuer:     cfg.JWT.Issuer,
		AccessTTL:  cfg.JWT.AccessTTL,
		RefreshTTL: cfg.JWT.RefreshTTL,
	})
	if err != nil {
		return err
	}

	repos := &repository.Repositories{
		User:    userrepo.NewRepository(postgresPool),
		Session: sessionrepo.NewRepository(postgresPool),
		Token:   tokenrepo.NewRepository(redisClient),
		Post:    postrepo.NewRepository(postgresPool),
		File:    filerepo.NewRepository(postgresPool),
		Storage: storage.NewStorage(s3Client, cfg.S3.Bucket, cfg.S3.PresignTTL),
	}

	services := &service.Services{
		Auth: authservice.NewService(repos.User, repos.Session, repos.Token, jwtManager),
		User: userservice.NewService(repos.User),
		Post: postservice.NewService(repos.Post, repos.File, repos.Storage),
		File: fileservice.NewService(repos.File, repos.Storage, cfg.S3.MaxUploadBytes),
	}

	cleanupService := cleanupservice.NewService(repos.File, repos.Storage, repos.Session, cfg.Cleanup.UnattachedTTL)
	cleanupCron := cron.NewCron(ctx, cleanupService, cfg.Cleanup)
	if err := cleanupCron.Start(); err != nil {
		return err
	}
	// Останавливаем cron раньше, чем закроются Postgres и Redis: defer выполняются в обратном порядке.
	defer func() {
		stopCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
		defer cancel()
		cleanupCron.Stop(stopCtx)
	}()

	server := &http.Server{
		Addr: ":" + cfg.HTTP.Port,
		Handler: handler.NewRouter(services, v1.Config{
			MaxUploadBytes: cfg.S3.MaxUploadBytes,
		}),
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	return serve(ctx, server, cfg.HTTP.ShutdownTimeout)
}

func serve(ctx context.Context, server *http.Server, shutdownTimeout time.Duration) error {
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("server started", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("listen and serve: %w", err)
	case <-ctx.Done():
	}

	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}
	return nil
}
