package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func New(ctx context.Context, cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		pingErr := fmt.Errorf("ping redis: %w", err)
		if err := client.Close(); err != nil {
			return nil, errors.Join(pingErr, fmt.Errorf("close redis: %w", err))
		}
		return nil, pingErr
	}
	return client, nil
}
