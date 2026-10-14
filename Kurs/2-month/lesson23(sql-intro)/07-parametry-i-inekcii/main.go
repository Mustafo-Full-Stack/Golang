package main

import (
	"database/sql"
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

// readUsers - тот же шаблон из шага 5 (Next -> Scan -> Err),
// вынесенный в функцию, чтобы не писать его в каждой функции заново.
func readUsers(rows *sql.Rows) ([]User, error) {
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// ПЛОХО! НИКОГДА ТАК НЕ ДЕЛАЙТЕ!
// Ввод пользователя склеивается прямо в текст SQL-запроса.
func findByNameUnsafe(db *sql.DB, name string) ([]User, error) {
	query := "SELECT id, name, email, age, is_active, created_at FROM users WHERE name = '" + name + "'"
	fmt.Println("  Запрос к базе:", query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return readUsers(rows)
}

// ХОРОШО: значение передаётся отдельно через плейсхолдер $1.
// Драйвер отправляет текст запроса и данные раздельно - данные никогда не станут кодом.
func findByNameSafe(db *sql.DB, name string) ([]User, error) {
	query := "SELECT id, name, email, age, is_active, created_at FROM users WHERE name = $1"
	fmt.Println("  Запрос к базе:", query, "| $1 =", name)

	rows, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	return readUsers(rows)
}

// Ещё один пример с параметром: пользователи не младше minAge
func findByAge(db *sql.DB, minAge int) ([]User, error) {
	rows, err := db.Query(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		WHERE age >= $1
		ORDER BY age
	`, minAge)
	if err != nil {
		return nil, err
	}
	return readUsers(rows)
}

func main() {
	db := mustConnect()
	defer db.Close()

	// Обычный пользователь вводит своё имя
	fmt.Println("=== Обычный ввод: Али ===")
	users, _ := findByNameUnsafe(db, "Али")
	fmt.Println("  Найдено:", len(users))
	users, _ = findByNameSafe(db, "Али")
	fmt.Println("  Найдено:", len(users))

	// А злоумышленник вводит вот такую строку
	hack := "' OR '1'='1"

	fmt.Println("\n=== Злой ввод: ' OR '1'='1 ===")
	fmt.Println("Небезопасный вариант:")
	users, _ = findByNameUnsafe(db, hack)
	fmt.Println("  Найдено:", len(users), "- вернулись ВСЕ пользователи, фильтр обойдён!")

	fmt.Println("Безопасный вариант:")
	users, _ = findByNameSafe(db, hack)
	fmt.Println("  Найдено:", len(users), "- ищется имя, которое буквально равно этой строке")

	// А если ввести '; DROP TABLE users; --  - таблицы просто не станет.
	// Это SQL-инъекция - уязвимость номер один в мире.
	// (Здесь мы это не запускаем, чтобы не потерять таблицу.)

	fmt.Println("\n=== Параметр $1: пользователи от 26 лет ===")
	users, err := findByAge(db, 26)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(" ", u.Name, u.Age)
	}

	// В PostgreSQL плейсхолдеры $1, $2, а в MySQL - ?
	// Правило: никогда не склеивайте SQL строками!
}
