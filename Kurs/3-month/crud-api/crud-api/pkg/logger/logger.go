package logger

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
)

type Config struct {
	Env   string
	Level slog.Level
	File  string
}

func New(cfg Config) (*slog.Logger, func() error, error) {
	opts := &slog.HandlerOptions{Level: cfg.Level}

	if cfg.Env != EnvLocal {
		return slog.New(slog.NewJSONHandler(os.Stdout, opts)), func() error { return nil }, nil
	}

	f, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}
	return slog.New(slog.NewTextHandler(f, opts)), f.Close, nil
}
