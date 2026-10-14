package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

// Подключение вынесли в функцию, чтобы не повторять его в каждой программе.
// Приставка must - "обязательно получится, иначе программа завершится"
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

// getAllUsers читает много строк - значит, используем db.Query.
// Этот шаблон из 5 пунктов повторяется ВСЕГДА, когда читаем несколько строк.
func getAllUsers(db *sql.DB) ([]User, error) {
	// 1. Делаем запрос и получаем rows
	rows, err := db.Query(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("запрос пользователей: %w", err)
	}

	// 2. СРАЗУ ставим defer rows.Close() - иначе соединение не вернётся в пул
	defer rows.Close()

	var users []User

	// 3. Идём по строкам
	for rows.Next() {
		var u User

		// 4. Scan: порядок полей СТРОГО как в SELECT
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("сканирование строки: %w", err)
		}
		users = append(users, u)
	}

	// 5. После цикла проверяем rows.Err() - часто забываемая деталь
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("обход результатов: %w", err)
	}

	return users, nil
}

func printUsers(users []User) {
	fmt.Printf("%-4s %-10s %-20s %-5s %s\n", "ID", "Имя", "Email", "Возр", "Активен")
	fmt.Println(strings.Repeat("-", 55))
	for _, u := range users {
		fmt.Printf("%-4d %-10s %-20s %-5d %t\n", u.ID, u.Name, u.Email, u.Age, u.IsActive)
	}
}

func main() {
	db := mustConnect()
	defer db.Close()

	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	printUsers(users)
}
