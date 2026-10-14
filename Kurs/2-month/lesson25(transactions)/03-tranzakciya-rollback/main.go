package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Переключатель для демонстрации.
// true  - после списания "падаем": транзакция откатится, деньги останутся на месте.
// false - перевод проходит до конца.
const simulateCrash = true

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

// transfer - тот же перевод, но в транзакции: "всё или ничего"
func transfer(db *sql.DB, from, to int, amount float64) error {
	// BEGIN - начинаем транзакцию
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}

	// Rollback сработает, если мы не дошли до Commit (return по ошибке или "сбой").
	// После успешного Commit он ничего не делает. Это стандартная идиома Go.
	defer tx.Rollback()

	// ВАЖНО: внутри транзакции все запросы идут через tx, а НЕ через db!
	_, err = tx.Exec(`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from)
	if err != nil {
		return fmt.Errorf("списание: %w", err)
	}

	if simulateCrash {
		return fmt.Errorf("сервер упал") // сбой между списанием и начислением
	}

	_, err = tx.Exec(`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to)
	if err != nil {
		return fmt.Errorf("начисление: %w", err)
	}

	// COMMIT - сохраняем всё разом
	return tx.Commit()
}

func main() {
	db := mustConnect()
	defer db.Close()

	fmt.Println("До перевода:")
	printBalances(db)

	err := transfer(db, 1, 2, 100)
	fmt.Println("Результат:", err)

	fmt.Println("\nПосле перевода:")
	printBalances(db)
	// Балансы НЕ изменились. Деньги на месте - вот за это базы данных и любят.

	// Поставьте simulateCrash = false и запустите ещё раз:
	// деньги переехали от Али к Вере, а общая сумма не изменилась.
}
