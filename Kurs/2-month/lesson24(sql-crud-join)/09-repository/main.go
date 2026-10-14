package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Паттерн Repository: весь SQL живёт в одном месте (в методах UserRepository),
// а остальной код работает с понятными методами и не знает, что внутри SQL.
// Если завтра сменится база - править придётся только здесь.

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

func (r *UserRepository) Create(u User) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id
	`, u.Name, u.Email, u.Age).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.Create: %w", err)
	}
	return id, nil
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	var u User
	err := r.db.QueryRow(`
		SELECT id, name, email, age FROM users WHERE id = $1
	`, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("repo.GetByID: пользователь %d не найден", id)
	}
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}
	return &u, nil
}

// List возвращает не больше limit пользователей, пропустив первые offset
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

func (r *UserRepository) Update(u User) error {
	res, err := r.db.Exec(`
		UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4
	`, u.Name, u.Email, u.Age, u.ID)
	if err != nil {
		return fmt.Errorf("repo.Update: %w", err)
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("repo.Update: пользователь %d не найден", u.ID)
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("repo.Delete: %w", err)
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("repo.Delete: пользователь %d не найден", id)
	}
	return nil
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

// Обратите внимание: в main нет ни одного слова SQL
func main() {
	db := mustConnect()
	defer db.Close()

	repo := NewUserRepository(db)

	id, err := repo.Create(User{Name: "Мадина", Email: "madina@mail.tj", Age: 22})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создан:", id)

	u, err := repo.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Получен: %+v\n", *u)

	u.Age = 23
	if err := repo.Update(*u); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Возраст обновлён")

	users, err := repo.List(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Всего в выборке:", len(users))

	if err := repo.Delete(id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Удалён")

	// Ошибки репозитория - понятные и с контекстом
	_, err = repo.GetByID(999)
	fmt.Println("Ошибка:", err)
}
