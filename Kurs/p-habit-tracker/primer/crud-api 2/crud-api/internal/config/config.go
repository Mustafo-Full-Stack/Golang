package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"crud-api/pkg/db/postgres"
	"crud-api/pkg/logger"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerConfigs ServerConfigs
	JWT           JWTConfig
	Logger        logger.Config
	Postgres      postgres.Config
}

type ServerConfigs struct {
	Port        string
	Environment string
}

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func Load() (*Config, error) {
	environmentName := getEnv("ENVIRONMENT", "prod")
	port := getEnv("HTTP_PORT", "8080")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	accessTTL, err := getEnvPositiveDuration("ACCESS_TOKEN_TTL", 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshTTL, err := getEnvPositiveDuration("REFRESH_TOKEN_TTL", 30*24*time.Hour)
	if err != nil {
		return nil, err
	}

	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(getEnv("LOG_LEVEL", "info"))); err != nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL: %w", err)
	}

	return &Config{
		ServerConfigs: ServerConfigs{
			Port:        port,
			Environment: environmentName,
		},
		JWT: JWTConfig{
			Secret:     jwtSecret,
			AccessTTL:  accessTTL,
			RefreshTTL: refreshTTL,
		},
		Logger: logger.Config{
			Env:   environmentName,
			Level: logLevel,
		},
		Postgres: postgres.Config{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			Database: getEnv("POSTGRES_DB", "crud_api"),
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

func getEnvDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := getEnv(key, defaultValue.String())

	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}

	return duration, nil
}

func getEnvPositiveDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	duration, err := getEnvDuration(key, defaultValue)
	if err != nil {
		return 0, err
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid %s: must be positive", key)
	}

	return duration, nil
}
