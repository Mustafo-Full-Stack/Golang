package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// Мини-проект: консольный просмотрщик пользователей.
// Соединяет меню (урок про циклы и switch) с базой данных.

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

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("запрос пользователей: %w", err)
	}
	return readUsers(rows)
}

func getUserByID(db *sql.DB, id int) (*User, error) {
	var u User
	err := db.QueryRow(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive, &u.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("пользователь с id=%d не найден", id)
	}
	if err != nil {
		return nil, fmt.Errorf("запрос пользователя: %w", err)
	}
	return &u, nil
}

func findByAge(db *sql.DB, minAge int) ([]User, error) {
	rows, err := db.Query(`
		SELECT id, name, email, age, is_active, created_at
		FROM users
		WHERE age >= $1
		ORDER BY age
	`, minAge)
	if err != nil {
		return nil, fmt.Errorf("поиск по возрасту: %w", err)
	}
	return readUsers(rows)
}

func getStats(db *sql.DB) error {
	var count int
	var avgAge, lowest, highest sql.NullFloat64

	err := db.QueryRow(`
		SELECT COUNT(*), AVG(age), MIN(age), MAX(age) FROM users
	`).Scan(&count, &avgAge, &lowest, &highest)
	if err != nil {
		return fmt.Errorf("статистика: %w", err)
	}

	fmt.Println("Всего пользователей:", count)
	if avgAge.Valid {
		fmt.Printf("Средний возраст: %.1f\n", avgAge.Float64)
		fmt.Printf("Диапазон: %.0f - %.0f\n", lowest.Float64, highest.Float64)
	} else {
		fmt.Println("Нет данных о возрасте")
	}
	return nil
}

func printUsers(users []User) {
	if len(users) == 0 {
		fmt.Println("Ничего не найдено")
		return
	}
	fmt.Printf("%-4s %-10s %-20s %-5s %s\n", "ID", "Имя", "Email", "Возр", "Активен")
	fmt.Println(strings.Repeat("-", 55))
	for _, u := range users {
		fmt.Printf("%-4d %-10s %-20s %-5d %t\n", u.ID, u.Name, u.Email, u.Age, u.IsActive)
	}
}

func main() {
	db := mustConnect()
	defer db.Close()

	for {
		fmt.Println("\n1 - все пользователи")
		fmt.Println("2 - найти по ID")
		fmt.Println("3 - фильтр по возрасту")
		fmt.Println("4 - статистика")
		fmt.Println("0 - выход")
		fmt.Print("Выбор: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			users, err := getAllUsers(db)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			printUsers(users)

		case 2:
			var id int
			fmt.Print("ID: ")
			fmt.Scanln(&id)

			u, err := getUserByID(db, id)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Printf("ID: %d\nИмя: %s\nEmail: %s\nВозраст: %d\nАктивен: %t\nСоздан: %s\n",
				u.ID, u.Name, u.Email, u.Age, u.IsActive, u.CreatedAt.Format("02.01.2006 15:04"))

		case 3:
			var minAge int
			fmt.Print("Минимальный возраст: ")
			fmt.Scanln(&minAge)

			users, err := findByAge(db, minAge)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			printUsers(users)

		case 4:
			if err := getStats(db); err != nil {
				fmt.Println("Ошибка:", err)
			}

		case 0:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Нет такого пункта меню")
		}
	}
}
