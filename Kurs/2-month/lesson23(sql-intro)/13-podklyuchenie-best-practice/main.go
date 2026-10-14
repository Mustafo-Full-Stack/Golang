package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Шаг 13. Как подключаются к базе в настоящих проектах.
// В шагах 3 и 4 строка подключения с паролем лежала прямо в коде - для урока это удобно,
// но в реальном проекте так делать нельзя: код попадает в git, и пароль видят все.
//
// Шаг состоит из нескольких файлов одного пакета:
//   main.go   - запуск: подключились, сделали запросы, запустили демонстрации
//   config.go - структура DBConfig, настройки из переменных окружения, сборка адреса
//   pool.go   - newPool: создание пула и проверка, что база отвечает
//   demo.go   - как ведёт себя пул: под нагрузкой и "живой"
//
// Что делаем правильно:
//   1. Все параметры подключения собираем в структуру DBConfig.
//   2. Значения берём из переменных окружения, а не из кода.
//   3. Строку подключения собирает функция dsn() через net/url: спецсимволы в пароле
//      (@, /, #, %) экранируются сами, склеивать строку руками не нужно.
//   4. Подключаемся с таймаутом и несколькими попытками: база может ещё запускаться.
//   5. Размер пула задаём в конфиге.
//   6. Пароль никогда не печатаем в логи.
//   7. log.Fatal вызываем только в main - так defer pool.Close() успевает сработать.
//   8. Программа с пулом живёт долго и завершается по Ctrl+C: пул закрываем аккуратно.
//
// Драйвер здесь pgx - в остальных шагах был lib/pq. pgx - современный драйвер для PostgreSQL,
// он сам умеет пул соединений (pgxpool), а контекст передаётся в каждый запрос.
//
// Запуск (база должна работать, см. docker-compose.yml).
// Файлов несколько, поэтому запускаем ПАПКУ, а не один файл (go run main.go не сработает):
//   go run ./13-podklyuchenie-best-practice
//   cd 13-podklyuchenie-best-practice && go run .
//   DB_PASSWORD=wrong go run ./13-podklyuchenie-best-practice

func main() {
	if err := run(); err != nil {
		log.Fatal("ошибка: ", err)
	}
}

func run() error {
	ctx := context.Background()
	cfg := loadConfig()

	// Печатаем всё, КРОМЕ пароля
	fmt.Printf("Подключаемся: %s@%s:%s, база %s\n", cfg.User, cfg.Host, cfg.Port, cfg.Database)

	pool, err := newPool(ctx, cfg)
	if err != nil {
		return err
	}
	defer pool.Close()

	fmt.Println("Подключение к базе установлено!")

	// Контекст в pgx обязателен в каждом запросе.
	// С таймаутом: зависший запрос не повесит программу
	queryCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var version string
	if err := pool.QueryRow(queryCtx, "SELECT version()").Scan(&version); err != nil {
		return fmt.Errorf("запрос версии: %w", err)
	}
	fmt.Println("Версия:", version)

	var count int
	if err := pool.QueryRow(queryCtx, "SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		return fmt.Errorf("подсчёт пользователей (выполните reset.sql, если нет таблицы users): %w", err)
	}
	fmt.Println("Пользователей в таблице:", count)

	// Часть 2: зачем нужен пул - много запросов одновременно (demo.go)
	demoPoolLoad(ctx, pool)

	// Часть 3: держим пул открытым, пока не нажмём Ctrl+C (demo.go).
	// Ctrl+C превращаем в отмену контекста: цикл в watchPool завершится сам,
	// и defer pool.Close() успеет закрыть пул.
	watchCtx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	watchPool(watchCtx, pool)

	// Эксперименты:
	// 1) DB_PASSWORD=wrong - неверный пароль: ошибка понятная, самого пароля в ней нет
	// 2) DB_PASSWORD='p@ss/word#1' - спецсимволы не ломают адрес (ошибка будет про пароль, а не про формат)
	// 3) остановите базу (docker compose stop) и запустите программу - увидите несколько попыток подряд
	// 4) уберите запасное значение пароля в loadConfig: почему в проде его быть не должно?
	// 5) в loadConfig поставьте MaxConns: 20 - часть 2 пройдёт в разы быстрее. Почему?
	// 6) пока работает часть 3: docker compose stop, потом docker compose start - пул переподключится сам
	return nil
}
