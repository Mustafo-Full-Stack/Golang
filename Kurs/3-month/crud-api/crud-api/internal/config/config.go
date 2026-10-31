package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"crud-api/pkg/db/postgres"
	"crud-api/pkg/db/redis"
	"crud-api/pkg/logger"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerConfigs   ServerConfigs
	IsRepoInMemory  bool
	IsTokenInMemory bool
	TokenTTL        time.Duration
	IsSlogEnabled   bool
	Logger          logger.Config
	Postgres        postgres.Config
	Redis           redis.Config
}

type ServerConfigs struct {
	Port        string
	Environment string
}

func Load() (*Config, error) {
	environmentName := getEnv("ENVIRONMENT", "prod")
	port := getEnv("HTTP_PORT", "8080")

	isRepoInMemory, err := getEnvBool("IS_REPO_IN_MEMORY", true)
	if err != nil {
		return nil, err
	}

	isTokenInMemory, err := getEnvBool("IS_TOKEN_IN_MEMORY", true)
	if err != nil {
		return nil, err
	}

	tokenTTL, err := getEnvDuration("TOKEN_TTL", 24*time.Hour)
	if err != nil {
		return nil, err
	}
	if tokenTTL <= 0 {
		return nil, fmt.Errorf("invalid TOKEN_TTL: must be positive")
	}

	isSlogEnabled, err := getEnvBool("IS_SLOG_ENABLED", false)
	if err != nil {
		return nil, err
	}

	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(getEnv("LOG_LEVEL", "info"))); err != nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL: %w", err)
	}

	redisDB, err := getEnvInt("REDIS_DB", 0)
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerConfigs: ServerConfigs{
			Port:        port,
			Environment: environmentName,
		},
		IsRepoInMemory:  isRepoInMemory,
		IsTokenInMemory: isTokenInMemory,
		TokenTTL:        tokenTTL,
		IsSlogEnabled:   isSlogEnabled,
		Logger: logger.Config{
			Env:   environmentName,
			Level: logLevel,
			File:  getEnv("LOG_FILE", "app.log"),
		},
		Postgres: postgres.Config{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			Database: getEnv("POSTGRES_DB", "crud_api"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		Redis: redis.Config{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) (int, error) {
	value := getEnv(key, strconv.Itoa(defaultValue))

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}

	return intValue, nil
}

func getEnvBool(key string, defaultValue bool) (bool, error) {
	value := getEnv(key, strconv.FormatBool(defaultValue))

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("invalid %s: %w", key, err)
	}

	return boolValue, nil
}

func getEnvDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := getEnv(key, defaultValue.String())

	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}

	return duration, nil
}
