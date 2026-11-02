package config

import (
	"fmt"
	"os"
	"time"

	"habit-tracker/pkg/db/postgres"
)

type Config struct {
	HTTPPort  string
	JWTSecret string
	JWTTTL    time.Duration
	Postgres  postgres.Config
}

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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
