package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// Типичные ошибки при работе с транзакциями.
// Код здесь нарочно "неправильный" - мы смотрим, как ошибка проявляется.
// Не исправляйте его, ошибка и есть урок.

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

func balanceOf(db *sql.DB, id int) float64 {
	var balance float64
	err := db.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, id).Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

// Ошибка 1: забыт Commit
func mistakeForgotCommit(db *sql.DB) {
	fmt.Println("--- 1. Забыт Commit ---")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Exec(`UPDATE accounts SET balance = 9999 WHERE id = 3`)
	// Commit() нет!

	fmt.Println("Баланс Тимура снаружи:", balanceOf(db, 3), "- не 9999, изменения не сохранены")
	fmt.Println("Соединений занято (транзакция висит):", db.Stats().InUse)

	tx.Rollback() // в демонстрации убираем за собой
}

// Ошибка 2: запрос через db вместо tx
func mistakeDBInsteadOfTx(db *sql.DB) {
	fmt.Println("\n--- 2. Запрос через db вместо tx ---")

	before := balanceOf(db, 1)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`UPDATE accounts SET balance = balance - 100 WHERE id = 1`) // мимо транзакции!
	tx.Rollback()

	after := balanceOf(db, 1)
	fmt.Printf("Баланс Али до: %.2f, после Rollback: %.2f - откат не помог!\n", before, after)

	db.Exec(`UPDATE accounts SET balance = balance + 100 WHERE id = 1`) // возвращаем
}

// Ошибка 3: Commit после Rollback
func mistakeCommitAfterRollback(db *sql.DB) {
	fmt.Println("\n--- 3. Commit после Rollback ---")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Rollback()

	err = tx.Commit()
	fmt.Println("Ошибка:", err)
}

// Ошибка 4: нарушение CHECK внутри транзакции
func mistakeCheckViolation(db *sql.DB) {
	fmt.Println("\n--- 4. Нарушение CHECK внутри транзакции ---")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE accounts SET balance = -500 WHERE id = 1`)
	fmt.Println("Ошибка:", err)
	// Хорошо: даже при ошибке в коде база защищает данные
}

// lockTwo - транзакция, которая блокирует два счёта по очереди
func lockTwo(db *sql.DB, name string, first, second int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`SELECT id FROM accounts WHERE id = $1 FOR UPDATE`, first)
	if err != nil {
		return fmt.Errorf("транзакция %s: %w", name, err)
	}
	fmt.Printf("  %s заблокировала счёт %d\n", name, first)

	time.Sleep(300 * time.Millisecond) // даём второй транзакции успеть захватить свой счёт

	_, err = tx.Exec(`SELECT id FROM accounts WHERE id = $1 FOR UPDATE`, second)
	if err != nil {
		return fmt.Errorf("транзакция %s: %w", name, err)
	}
	fmt.Printf("  %s заблокировала счёт %d\n", name, second)

	return tx.Commit()
}

// runTwoTransactions запускает две транзакции одновременно
func runTwoTransactions(db *sql.DB, a1, a2, b1, b2 int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := lockTwo(db, "A", a1, a2); err != nil {
			fmt.Println("  ✗", err)
		} else {
			fmt.Println("  ✓ A завершилась")
		}
	}()
	go func() {
		defer wg.Done()
		if err := lockTwo(db, "B", b1, b2); err != nil {
			fmt.Println("  ✗", err)
		} else {
			fmt.Println("  ✓ B завершилась")
		}
	}()

	wg.Wait()
}

// Ошибка 5: deadlock - две транзакции ждут друг друга
func mistakeDeadlock(db *sql.DB) {
	fmt.Println("\n--- 5. Deadlock: блокируем счета в РАЗНОМ порядке ---")
	fmt.Println("A: сначала 1, потом 2.  B: сначала 2, потом 1.")
	runTwoTransactions(db, 1, 2, 2, 1)
	// База сама обнаруживает взаимную блокировку (примерно за секунду)
	// и откатывает одну из транзакций: deadlock detected

	fmt.Println("\n--- Решение: блокируем в ОДНОМ порядке (по возрастанию id) ---")
	fmt.Println("A: сначала 1, потом 2.  B: тоже сначала 1, потом 2.")
	runTwoTransactions(db, 1, 2, 1, 2)
	// B просто дождётся, пока A закончит
}

// Ещё одна ошибка - долгая транзакция.
// Если транзакция висит открытой минуту, все остальные, кому нужны те же строки, ждут.
// Транзакции должны быть короткими: никаких HTTP-запросов и пауз внутри.

func main() {
	db := mustConnect()
	defer db.Close()

	mistakeForgotCommit(db)
	mistakeDBInsteadOfTx(db)
	mistakeCommitAfterRollback(db)
	mistakeCheckViolation(db)
	mistakeDeadlock(db)
}
