package main

import (
	"database/sql"
	"fmt"
	"log"

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

// withTransaction - универсальная обёртка. Вся рутина (Begin, Rollback, Commit)
// спрятана здесь, а сама работа передаётся как функция fn
// (функция как параметр - то, что мы проходили в уроке про функции).
func withTransaction(db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}

	// Если внутри fn случится panic - откатываем и пробрасываем панику дальше
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// fn вернула ошибку - откатываем всё
	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("ошибка %v, откат тоже не удался: %w", err, rbErr)
		}
		return err
	}

	// Ошибок нет - сохраняем
	return tx.Commit()
}

// Использование: в функции остаётся только бизнес-логика, без Begin/Commit/Rollback
func transferSimple(db *sql.DB, from, to int, amount float64) error {
	return withTransaction(db, func(tx *sql.Tx) error {
		var balance float64
		err := tx.QueryRow(
			`SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, from,
		).Scan(&balance)
		if err != nil {
			return fmt.Errorf("чтение счёта отправителя %d: %w", from, err)
		}
		if balance < amount {
			return fmt.Errorf("недостаточно средств")
		}

		_, err = tx.Exec(
			`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from,
		)
		if err != nil {
			return err
		}

		result, err := tx.Exec(
			`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to,
		)
		if err != nil {
			return err
		}

		// Если получателя нет, UPDATE не выдаст ошибку, а просто ничего не изменит.
		// Тогда деньги бы пропали, поэтому проверяем RowsAffected.
		if n, _ := result.RowsAffected(); n == 0 {
			return fmt.Errorf("счёт получателя %d не найден", to)
		}
		return nil
	})
}

func main() {
	db := mustConnect()
	defer db.Close()

	fmt.Println("До переводов:")
	printBalances(db)

	fmt.Println("\n--- Перевод 100 от Веры к Тимуру ---")
	if err := transferSimple(db, 2, 3, 100); err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Готово")
	}

	fmt.Println("\n--- Недостаточно средств ---")
	if err := transferSimple(db, 3, 1, 99999); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Несуществующий получатель (деньги уже списаны, но откат вернёт их) ---")
	if err := transferSimple(db, 1, 999, 50); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\nПосле переводов:")
	printBalances(db)
}
