package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func mustConnect() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=coursedb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("не удалось создать пул:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("не удалось подключиться:", err)
	}
	return db
}

func printBalances(db *sql.DB) {
	rows, err := db.Query(`SELECT id, owner, balance FROM accounts ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	total := 0.0
	for rows.Next() {
		var id int
		var owner string
		var balance float64
		if err := rows.Scan(&id, &owner, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  %d. %-8s %8.2f\n", id, owner, balance)
		total += balance
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Всего денег в системе: %.2f\n", total)
}

// transferWithTimeout - перевод с ограничением по времени.
// context пронизывает всю экосистему Go, в том числе работу с БД (урок про context):
// BeginTx, ExecContext, QueryContext - те же методы, но с ctx первым параметром.
func transferWithTimeout(ctx context.Context, db *sql.DB, from, to int, amount float64, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil) // nil - настройки по умолчанию
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx,
		`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from)
	if err != nil {
		return fmt.Errorf("списание: %w", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("счёт отправителя %d не найден", from)
	}

	// Имитация медленной базы: этот запрос "думает" 2 секунды
	_, err = tx.ExecContext(ctx, `SELECT pg_sleep(2)`)
	if err != nil {
		return fmt.Errorf("долгий запрос: %w", err)
	}

	result, err = tx.ExecContext(ctx,
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to)
	if err != nil {
		return fmt.Errorf("начисление: %w", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("счёт получателя %d не найден", to)
	}

	return tx.Commit()
}

func main() {
	db := mustConnect()
	defer db.Close()

	ctx := context.Background()

	fmt.Println("До переводов:")
	printBalances(db)

	fmt.Println("\n--- Таймаут 5 секунд, база думает 2 секунды ---")
	err := transferWithTimeout(ctx, db, 1, 2, 100, 5*time.Second)
	if err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Перевод выполнен")
	}

	fmt.Println("\n--- Таймаут 1 секунда, база думает 2 секунды ---")
	err = transferWithTimeout(ctx, db, 1, 2, 100, 1*time.Second)
	if err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Перевод выполнен")
	}

	// Первый перевод прошёл (Али -100, Вера +100). Второй отменён по таймауту,
	// и транзакция откатилась сама: списание второго перевода не сохранилось.
	fmt.Println("\nПосле переводов:")
	printBalances(db)
}
