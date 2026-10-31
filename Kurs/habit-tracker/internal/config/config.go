package config

import (
	"fmt"
	"os"
	"time"

	"habit-tracker/pkg/db/postgres"
)

// Config хранит настройки приложения из переменных окружения.
type Config struct {
	HTTPPort  string
	JWTSecret string
	JWTTTL    time.Duration
	Postgres  postgres.Config
}

// Load читает переменные окружения и проверяет обязательные настройки.
func Load() (*Config, error) {
	ttl, err := time.ParseDuration(getEnv("JWT_TTL", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_TTL: %w", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return &Config{
		HTTPPort:  getEnv("HTTP_PORT", "8080"),
		JWTSecret: jwtSecret,
		JWTTTL:    ttl,
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

// getEnv возвращает значение переменной или значение по умолчанию.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
