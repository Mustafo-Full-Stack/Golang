package cron

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"

	"demo-project/internal/config"
)

type cleanupService interface {
	DeleteUnattachedFiles(ctx context.Context) (int, error)
	DeleteExpiredSessions(ctx context.Context) (int64, error)
}

type Cron struct {
	scheduler *gocron.Scheduler
	service   cleanupService
	cfg       config.Cleanup
	ctx       context.Context
}

func NewCron(ctx context.Context, service cleanupService, cfg config.Cleanup) *Cron {
	return &Cron{
		scheduler: gocron.NewScheduler(time.UTC),
		service:   service,
		cfg:       cfg,
		ctx:       ctx,
	}
}

func (c *Cron) Start() error {
	if !c.cfg.Enabled {
		slog.Info("cleanup cron is disabled")
		return nil
	}

	// SingletonMode не даёт запускам накладываться друг на друга.
	_, err := c.scheduler.Every(c.cfg.Interval).SingletonMode().Do(c.cleanup)
	if err != nil {
		return fmt.Errorf("schedule cleanup: %w", err)
	}

	c.scheduler.StartAsync()
	slog.Info("cleanup cron started", "interval", c.cfg.Interval)
	return nil
}

// Stop ждёт текущий проход, но не дольше, чем живёт ctx.
func (c *Cron) Stop(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		c.scheduler.Stop()
		close(done)
	}()

	select {
	case <-done:
		slog.InfoContext(ctx, "cleanup cron stopped")
	case <-ctx.Done():
		slog.WarnContext(ctx, "cleanup cron stop timed out", "error", ctx.Err())
	}
}

func (c *Cron) cleanup() {
	files, err := c.service.DeleteUnattachedFiles(c.ctx)
	if err != nil {
		slog.ErrorContext(c.ctx, "cleanup: delete unattached files failed", "error", err)
	} else {
		slog.InfoContext(c.ctx, "cleanup: unattached files deleted", "count", files)
	}

	sessions, err := c.service.DeleteExpiredSessions(c.ctx)
	if err != nil {
		slog.ErrorContext(c.ctx, "cleanup: delete expired sessions failed", "error", err)
	} else {
		slog.InfoContext(c.ctx, "cleanup: expired sessions deleted", "count", sessions)
	}
}
