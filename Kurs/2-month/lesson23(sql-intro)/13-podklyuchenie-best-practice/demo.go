package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// printStat печатает, что сейчас происходит внутри пула
func printStat(pool *pgxpool.Pool) {
	s := pool.Stat()
	fmt.Printf("  пул: занято %d, свободно %d, всего %d из %d\n",
		s.AcquiredConns(), s.IdleConns(), s.TotalConns(), s.MaxConns())
}

// demoPoolLoad - часть 2. Зачем нужен пул.
// Запускаем 20 запросов одновременно, а в пуле всего несколько соединений (MaxConns).
// Остальные запросы ждут в очереди, пока соединение освободится.
func demoPoolLoad(ctx context.Context, pool *pgxpool.Pool) {
	const queries = 20

	fmt.Printf("\n--- Часть 2. %d запросов одновременно, соединений в пуле: %d ---\n",
		queries, pool.Stat().MaxConns())

	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < queries; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// pg_sleep(0.5) - запрос, который "думает" полсекунды. Результат нам не нужен - Exec
			if _, err := pool.Exec(ctx, "SELECT pg_sleep(0.5)"); err != nil {
				fmt.Println("запрос не удался:", err)
			}
		}()
	}

	// Пока запросы работают, раз в 300 мс смотрим на пул
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Printf("Все запросы выполнены за %v\n", time.Since(start).Round(100*time.Millisecond))
			printStat(pool) // соединения не закрылись: лежат в пуле и ждут следующих запросов
			return
		case <-ticker.C:
			printStat(pool)
		}
	}
}

// watchPool - часть 3. Держим пул открытым, пока контекст не отменят (Ctrl+C).
// Раз в 2 секунды проверяем базу. Остановите базу (docker compose stop), а потом запустите
// снова (docker compose start): пул сам переподключится, программу перезапускать не нужно.
func watchPool(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Println("\n--- Часть 3. Пул открыт. Каждые 2 секунды проверяем базу. Выход - Ctrl+C ---")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nПолучен Ctrl+C, закрываем пул и выходим")
			return
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, time.Second)
			err := pool.Ping(pingCtx)
			cancel()

			if err != nil {
				fmt.Println("база не отвечает")
			} else {
				fmt.Println("база отвечает")
			}
			printStat(pool)
		}
	}
}
