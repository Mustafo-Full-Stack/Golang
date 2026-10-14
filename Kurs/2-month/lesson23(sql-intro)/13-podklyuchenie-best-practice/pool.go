package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// newPool создаёт пул соединений и дожидается, пока база ответит.
// Если база недоступна - возвращает ошибку, а не завершает программу сама.
func newPool(ctx context.Context, cfg DBConfig) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.dsn())
	if err != nil {
		return nil, fmt.Errorf("разбор настроек подключения: %w", err)
	}
	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.ConnConfig.ConnectTimeout = cfg.ConnectTimeout

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("создание пула: %w", err)
	}

	// Создание пула ничего не открывает (как sql.Open), поэтому проверяем сами.
	// Пробуем несколько раз: в docker база стартует не мгновенно.
	for attempt := 1; attempt <= cfg.ConnectRetries; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, cfg.ConnectTimeout)
		err = pool.Ping(pingCtx)
		cancel()
		if err == nil {
			return pool, nil
		}

		fmt.Printf("Попытка %d из %d: база не ответила\n", attempt, cfg.ConnectRetries)
		if attempt < cfg.ConnectRetries {
			time.Sleep(time.Second) // пауза перед следующей попыткой
		}
	}

	pool.Close() // пул уже создан - не забываем закрыть, раз возвращаем ошибку
	return nil, fmt.Errorf("база не отвечает после %d попыток: %w", cfg.ConnectRetries, err)
}
