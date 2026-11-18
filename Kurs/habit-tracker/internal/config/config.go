package config

import (
	"fmt"
	"os"
	"time"

	"habit-tracker/pkg/db/postgres"
)

type Config struct {
	HTTPPort   string
	JWTSecret  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Postgres   postgres.Config
}

func Load() (*Config, error) {
	accessTTL, err := time.ParseDuration(getEnv("ACCESS_TOKEN_TTL", "15m"))
	if err != nil || accessTTL <= 0 {
		return nil, fmt.Errorf("invalid ACCESS_TOKEN_TTL: %w", err)
	}

	refreshTTL, err := time.ParseDuration(getEnv("REFRESH_TOKEN_TTL", "720h"))
	if err != nil || refreshTTL <= 0 {
		return nil, fmt.Errorf("invalid REFRESH_TOKEN_TTL: %w", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return &Config{
		HTTPPort:   getEnv("HTTP_PORT", "8080"),
		JWTSecret:  jwtSecret,
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
		Postgres: postgres.Config{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			Database: getEnv("POSTGRES_DB", "habits"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
