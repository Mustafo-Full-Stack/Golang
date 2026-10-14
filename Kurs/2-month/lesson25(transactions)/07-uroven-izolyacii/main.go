package main

import (
	"context"
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

// readTwice открывает транзакцию с заданным уровнем изоляции и читает баланс Али ДВАЖДЫ.
// Между чтениями кто-то другой (вне нашей транзакции) пополняет его счёт на 100.
// Вопрос: увидим ли мы это изменение внутри своей транзакции?
func readTwice(db *sql.DB, level sql.IsolationLevel, name string) error {
	ctx := context.Background()

	// TxOptions - настройки транзакции: уровень изоляции и режим "только чтение"
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: level})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var first, second float64

	err = tx.QueryRow(`SELECT balance FROM accounts WHERE id = 1`).Scan(&first)
	if err != nil {
		return err
	}

	// "Другой пользователь" пополняет счёт - через db, то есть вне нашей транзакции
	_, err = db.Exec(`UPDATE accounts SET balance = balance + 100 WHERE id = 1`)
	if err != nil {
		return err
	}

	err = tx.QueryRow(`SELECT balance FROM accounts WHERE id = 1`).Scan(&second)
	if err != nil {
		return err
	}

	fmt.Printf("%-16s первое чтение: %.2f, второе чтение: %.2f\n", name, first, second)
	return tx.Commit()
}

// tryWriteInReadOnly: транзакция "только чтение" (например, для отчётов)
// не даёт ничего менять, а база может оптимизировать такие транзакции.
func tryWriteInReadOnly(db *sql.DB) {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE accounts SET balance = 0 WHERE id = 1`)
	fmt.Println("Ошибка:", err)
}

func main() {
	db := mustConnect()
	defer db.Close()

	// По умолчанию в PostgreSQL - READ COMMITTED
	fmt.Println("=== Уровни изоляции ===")
	readTwice(db, sql.LevelReadCommitted, "READ COMMITTED:")
	readTwice(db, sql.LevelRepeatableRead, "REPEATABLE READ:")
	readTwice(db, sql.LevelSerializable, "SERIALIZABLE:")

	// READ COMMITTED: второе чтение увидело чужое изменение ("неповторяющееся чтение").
	// REPEATABLE READ и SERIALIZABLE: внутри транзакции данные не меняются.
	// Чем строже уровень - тем безопаснее, но тем больше блокировок и медленнее.

	// Мы трижды пополнили счёт Али на 100 - вернём как было
	db.Exec(`UPDATE accounts SET balance = balance - 300 WHERE id = 1`)

	fmt.Println("\n=== Транзакция только для чтения ===")
	tryWriteInReadOnly(db)
}
