package main

import (
	"net"
	"net/url"
	"os"
	"time"
)

// DBConfig - всё, что нужно знать, чтобы подключиться к базе.
// Название без "PG": структура описывает любую базу, а не только PostgreSQL.
type DBConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	Database       string
	SSLMode        string
	MaxConns       int32         // максимум соединений в пуле
	ConnectRetries int           // сколько раз пробуем подключиться
	ConnectTimeout time.Duration // сколько ждём одну попытку
}

// в реальных проектах не должно быть в проде fallback для значений переменных окружений!
func loadConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "secret"),
		Database: getEnv("DB_NAME", "coursedb"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),

		// Пул специально маленький, чтобы в части 2 было видно очередь из запросов
		MaxConns:       5,
		ConnectRetries: 5,
		ConnectTimeout: 3 * time.Second,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// dsn собирает адрес подключения: postgres://user:password@host:port/database?sslmode=...
// url.UserPassword сам экранирует спецсимволы в логине и пароле.
func (c DBConfig) dsn() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     net.JoinHostPort(c.Host, c.Port),
		Path:     c.Database,
		RawQuery: url.Values{"sslmode": {c.SSLMode}}.Encode(),
	}
	return u.String()
}
