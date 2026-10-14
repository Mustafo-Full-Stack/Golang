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

// printBalances показывает все счета и сумму денег в системе.
// Эта сумма - главный показатель: перевод не должен её менять.
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

// transferBroken - перевод БЕЗ транзакции. Две операции живут сами по себе:
// первая выполнилась и сохранилась, вторая - не успела.
func transferBroken(db *sql.DB, from, to int, amount float64) error {
	// 1. Списываем
	_, err := db.Exec(`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from)
	if err != nil {
		return err
	}

	fmt.Printf("Деньги списаны со счёта %d, на счёт %d ещё не пришли... имитируем сбой!\n", from, to)
	return fmt.Errorf("сервер упал") // авария

	// 2. Начисление уже не выполнится!
	// _, err = db.Exec(`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to)
}

func main() {
	db := mustConnect()
	defer db.Close()

	fmt.Println("До перевода:")
	printBalances(db)

	err := transferBroken(db, 1, 2, 100)
	fmt.Println("Результат:", err)

	fmt.Println("\nПосле перевода:")
	printBalances(db)
	// У Али на 100 меньше, у Веры без изменений - 100 сомони ИСЧЕЗЛИ из системы.

	// После этого шага выполните reset.sql, чтобы вернуть деньги на место.
}
