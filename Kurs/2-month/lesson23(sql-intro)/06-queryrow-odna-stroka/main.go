package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Age       int
	IsActive  bool
	CreatedAt time.Time
}

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

// getUserByID ищет одного пользователя - значит, db.QueryRow.
// Здесь нет ни rows.Next(), ни rows.Close(): строка всего одна,
// и Scan вызывается сразу после QueryRow.
func getUserByID(db *sql.DB, id int) (*User, error) {
	var u User

	// $1 - место для параметра. Значение id передаём отдельно, после текста запроса
	err := db.QueryRow(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive, &u.CreatedAt)

	// sql.ErrNoRows - это не "всё сломалось", а "ничего не нашлось".
	// Это такая же sentinel-ошибка, как io.EOF (урок про ошибки), поэтому сравниваем через errors.Is
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("пользователь с id=%d не найден", id)
	}
	if err != nil {
		return nil, fmt.Errorf("запрос пользователя: %w", err)
	}

	return &u, nil
}

func main() {
	db := mustConnect()
	defer db.Close()

	// 1 - есть в базе, 999 - нет
	for _, id := range []int{1, 999} {
		u, err := getUserByID(db, id)
		if err != nil {
			fmt.Println("✗", err)
			continue
		}
		fmt.Printf("✓ %s (%s), %d лет\n", u.Name, u.Email, u.Age)
	}
}
