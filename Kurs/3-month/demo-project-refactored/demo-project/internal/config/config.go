package config

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App      App
	HTTP     HTTP
	Postgres Postgres
	Redis    Redis
	JWT      JWT
	S3       S3
	Cleanup  Cleanup
}

type App struct {
	Env      string     `envconfig:"APP_ENV" default:"local"` // local | prod
	LogLevel slog.Level `envconfig:"LOG_LEVEL" default:"info"`
	LogFile  string     `envconfig:"LOG_FILE" default:"app.log"` // только для APP_ENV=local
}

type HTTP struct {
	Port            string        `envconfig:"HTTP_PORT" default:"8080"`
	ReadTimeout     time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"30s"`
	WriteTimeout    time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
}

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	User     string `envconfig:"POSTGRES_USER" default:"demo"`
	Password string `envconfig:"POSTGRES_PASSWORD" default:"demo"`
	Database string `envconfig:"POSTGRES_DB" default:"demo-db"`
	SSLMode  string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
}

type Redis struct {
	Host     string `envconfig:"REDIS_HOST" default:"localhost"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
}

type JWT struct {
	Secret     string        `envconfig:"JWT_SECRET" required:"true"`
	Issuer     string        `envconfig:"JWT_ISSUER" default:"demo-project"`
	AccessTTL  time.Duration `envconfig:"JWT_ACCESS_TTL" default:"15m"`
	RefreshTTL time.Duration `envconfig:"JWT_REFRESH_TTL" default:"720h"`
}

type S3 struct {
	Endpoint       string        `envconfig:"S3_ENDPOINT" default:"localhost:9000"`
	AccessKey      string        `envconfig:"S3_ACCESS_KEY" default:"minioadmin"`
	SecretKey      string        `envconfig:"S3_SECRET_KEY" default:"minioadmin"`
	Bucket         string        `envconfig:"S3_BUCKET" default:"demo-media"`
	UseSSL         bool          `envconfig:"S3_USE_SSL" default:"false"`
	PresignTTL     time.Duration `envconfig:"S3_PRESIGN_TTL" default:"15m"`
	MaxUploadBytes int64         `envconfig:"S3_MAX_UPLOAD_BYTES" default:"10485760"`
}

type Cleanup struct {
	Enabled       bool          `envconfig:"CLEANUP_ENABLED" default:"true"`
	Interval      time.Duration `envconfig:"CLEANUP_INTERVAL" default:"30m"`
	UnattachedTTL time.Duration `envconfig:"CLEANUP_UNATTACHED_TTL" default:"24h"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) validate() error {
	if c.S3.PresignTTL <= 0 {
		return fmt.Errorf("S3_PRESIGN_TTL must be positive")
	}
	if c.S3.MaxUploadBytes <= 0 {
		return fmt.Errorf("S3_MAX_UPLOAD_BYTES must be positive")
	}
	if c.Cleanup.Interval <= 0 {
		return fmt.Errorf("CLEANUP_INTERVAL must be positive")
	}
	if c.Cleanup.UnattachedTTL <= 0 {
		return fmt.Errorf("CLEANUP_UNATTACHED_TTL must be positive")
	}
	return nil
}
