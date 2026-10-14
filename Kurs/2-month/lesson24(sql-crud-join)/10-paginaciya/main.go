package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) List(limit, offset int) ([]User, error) {
	rows, err := r.db.Query(`
		SELECT id, name, email, age FROM users ORDER BY id LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repo.List: %w", err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
			return nil, fmt.Errorf("repo.List: %w", err)
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// ListPaginated возвращает одну страницу пользователей и общее количество.
// page - номер страницы (с 1), perPage - сколько записей на странице.
func (r *UserRepository) ListPaginated(page, perPage int) ([]User, int, error) {
	// Сколько всего пользователей (чтобы знать, сколько будет страниц)
	var total int
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("repo.ListPaginated: %w", err)
	}

	// Страница 1 -> пропускаем 0 строк, страница 2 -> пропускаем perPage строк, и т.д.
	offset := (page - 1) * perPage

	users, err := r.List(perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
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

func main() {
	db := mustConnect()
	defer db.Close()

	repo := NewUserRepository(db)
	perPage := 3

	// Пока не знаем, сколько страниц - узнаем после первого запроса
	for page := 1; ; page++ {
		users, total, err := repo.ListPaginated(page, perPage)
		if err != nil {
			log.Fatal(err)
		}

		totalPages := (total + perPage - 1) / perPage // делим с округлением вверх
		fmt.Printf("=== Страница %d из %d (всего пользователей: %d) ===\n", page, totalPages, total)
		for _, u := range users {
			fmt.Printf("  %d. %s, %d лет\n", u.ID, u.Name, u.Age)
		}

		if page >= totalPages {
			break
		}
	}
}
