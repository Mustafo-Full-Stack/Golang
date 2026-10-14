package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq" // теперь без "_": нам нужен тип pq.Error, мы им пользуемся сами
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

func createUserReturningID(db *sql.DB, name, email string, age int) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`,
		name, email, age,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("вставка пользователя: %w", err)
	}
	return id, nil
}

// isUniqueViolation проверяет: это ошибка "значение уже существует"?
// Достаём ошибку драйвера через errors.As (как в уроке про ошибки)
// и смотрим на её код. У PostgreSQL для каждой ошибки есть свой код.
func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505" // 23505 = unique_violation
	}
	return false
}

func main() {
	db := mustConnect()
	defer db.Close()

	// Пытаемся создать пользователя с email, который уже занят (ali@mail.tj есть в базе)
	_, err := createUserReturningID(db, "Дубль", "ali@mail.tj", 20)
	if err != nil {
		// Сырая ошибка: непонятно, что с ней делать в коде
		fmt.Println("Сырая ошибка:", err)

		// Достаём из неё код ошибки PostgreSQL
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			fmt.Println("Код ошибки:", pqErr.Code, "-", pqErr.Code.Name())
		}

		// Теперь можно ответить пользователю по-человечески
		if isUniqueViolation(err) {
			fmt.Println("Такой email уже зарегистрирован")
		} else {
			fmt.Println("Неизвестная ошибка:", err)
		}
	}

	// Другая ошибка: нарушено ограничение CHECK (возраст не может быть отрицательным)
	_, err = createUserReturningID(db, "Плохой", "bad@mail.tj", -5)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			fmt.Println("\nКод ошибки:", pqErr.Code, "-", pqErr.Code.Name())
		}
		fmt.Println("Ошибка:", err)
	}
}
