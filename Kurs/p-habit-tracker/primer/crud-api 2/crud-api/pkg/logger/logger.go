package logger

import (
	"log/slog"
	"os"
)

const EnvLocal = "local"

type Config struct {
	Env   string
	Level slog.Level
}

// New пишет логи в stdout: локально — читаемым текстом, в остальных окружениях — JSON.
func New(cfg Config) *slog.Logger {
	opts := &slog.HandlerOptions{Level: cfg.Level}

	if cfg.Env == EnvLocal {
		return slog.New(slog.NewTextHandler(os.Stdout, opts))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
