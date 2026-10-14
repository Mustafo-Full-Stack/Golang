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

// createUser добавляет пользователя.
// INSERT ничего не возвращает (нет выборки), поэтому используем db.Exec, а не Query.
func createUser(db *sql.DB, name, email string, age int) error {
	result, err := db.Exec(
		`INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`,
		name, email, age,
	)
	if err != nil {
		return fmt.Errorf("вставка пользователя: %w", err)
	}

	// Exec возвращает sql.Result. RowsAffected - сколько строк затронул запрос
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Добавлено строк: %d\n", rows)
	return nil
}

func main() {
	db := mustConnect()
	defer db.Close()

	err := createUser(db, "Шерзод", "sherzod@mail.tj", 22)
	if err != nil {
		log.Fatal(err)
	}

	// Запустите программу второй раз - получите ошибку: такой email уже есть.
	// Разберём такие ошибки на шаге 3.
}
