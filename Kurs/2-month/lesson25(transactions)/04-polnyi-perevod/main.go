package main

import (
	"database/sql"
	"errors"
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

// transferMoney - полноценный перевод со всеми проверками.
// (Для учебного примера сумма хранится в float64. В реальных банках деньги
// считают в целых копейках или в специальном типе decimal.)
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	// Проверки, которым база не нужна, делаем ДО начала транзакции
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть положительной, получено %.2f", amount)
	}
	if fromID == toID {
		return fmt.Errorf("нельзя переводить самому себе")
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}
	defer tx.Rollback()

	// 1. Проверяем отправителя и блокируем его строку.
	// FOR UPDATE - никто другой не сможет изменить этот счёт, пока мы не закончим
	var fromBalance float64
	var fromOwner string
	err = tx.QueryRow(`
		SELECT owner, balance FROM accounts WHERE id = $1 FOR UPDATE
	`, fromID).Scan(&fromOwner, &fromBalance)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("счёт отправителя %d не найден", fromID)
	}
	if err != nil {
		return fmt.Errorf("чтение отправителя: %w", err)
	}

	if fromBalance < amount {
		return fmt.Errorf("недостаточно средств: на счёте %.2f, требуется %.2f", fromBalance, amount)
	}

	// 2. Проверяем получателя (тоже блокируем)
	var toOwner string
	err = tx.QueryRow(`SELECT owner FROM accounts WHERE id = $1 FOR UPDATE`, toID).Scan(&toOwner)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("счёт получателя %d не найден", toID)
	}
	if err != nil {
		return fmt.Errorf("чтение получателя: %w", err)
	}

	// 3. Списываем
	_, err = tx.Exec(`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, fromID)
	if err != nil {
		return fmt.Errorf("списание: %w", err)
	}

	// 4. Начисляем
	_, err = tx.Exec(`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, toID)
	if err != nil {
		return fmt.Errorf("начисление: %w", err)
	}

	// 5. Записываем историю
	_, err = tx.Exec(`INSERT INTO transfers (from_id, to_id, amount) VALUES ($1, $2, $3)`,
		fromID, toID, amount)
	if err != nil {
		return fmt.Errorf("запись истории: %w", err)
	}

	// 6. Подтверждаем всё разом
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("подтверждение транзакции: %w", err)
	}

	fmt.Printf("✓ Переведено %.2f: %s → %s\n", amount, fromOwner, toOwner)
	return nil
}

func main() {
	db := mustConnect()
	defer db.Close()

	printBalances(db)

	fmt.Println("\n--- Успешный перевод 300 ---")
	if err := transferMoney(db, 1, 2, 300); err != nil {
		fmt.Println("✗", err)
	}
	printBalances(db)

	fmt.Println("\n--- Перевод больше баланса ---")
	if err := transferMoney(db, 3, 1, 99999); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Несуществующий счёт ---")
	if err := transferMoney(db, 1, 999, 50); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Отрицательная сумма ---")
	if err := transferMoney(db, 1, 2, -100); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Итог (неудачные переводы ничего не изменили) ---")
	printBalances(db)
}
