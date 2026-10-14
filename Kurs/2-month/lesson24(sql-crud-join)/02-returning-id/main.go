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

// В PostgreSQL result.LastInsertId() не работает (покажем в "типичных ошибках").
// Вместо него добавляем в конец запроса RETURNING id -
// тогда INSERT возвращает строку с id, и мы читаем её как обычный QueryRow.
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

// RETURNING может вернуть любые поля, в том числе заполненные самой базой (id, created_at).
// Так можно получить созданную запись целиком.
func createUserFull(db *sql.DB, name, email string, age int) (User, error) {
	var u User
	err := db.QueryRow(`
		INSERT INTO users (name, email, age) VALUES ($1, $2, $3)
		RETURNING id, name, email, age, created_at
	`, name, email, age).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt)
	if err != nil {
		return User{}, fmt.Errorf("вставка пользователя: %w", err)
	}
	return u, nil
}

func main() {
	db := mustConnect()
	defer db.Close()

	id, err := createUserReturningID(db, "Фаррух", "farruh@mail.tj", 28)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создан пользователь с ID:", id)

	u, err := createUserFull(db, "Дилшод", "dilshod@mail.tj", 31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Создан: ID=%d, %s, создан в %s\n", u.ID, u.Name, u.CreatedAt.Format("15:04:05"))
}
