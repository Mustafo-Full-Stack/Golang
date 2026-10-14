package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
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

// updateUserAge меняет возраст одного пользователя.
// ВАЖНО: UPDATE несуществующей строки - НЕ ошибка. Запрос выполнился успешно,
// просто ничего не изменил. Поэтому обязательно проверяем RowsAffected.
func updateUserAge(db *sql.DB, id, newAge int) error {
	result, err := db.Exec(`UPDATE users SET age = $1 WHERE id = $2`, newAge, id)
	if err != nil {
		return fmt.Errorf("обновление возраста: %w", err)
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("пользователь с id=%d не найден", id)
	}
	fmt.Printf("Обновлено строк: %d\n", affected)
	return nil
}

// updateUser обновляет сразу несколько полей
func updateUser(db *sql.DB, u User) error {
	_, err := db.Exec(`
		UPDATE users SET name = $1, email = $2, age = $3
		WHERE id = $4
	`, u.Name, u.Email, u.Age, u.ID)
	return err
}

// patchUser обновляет только те поля, которые заданы (не nil).
// SQL собирается из частей, а значения идут через args - инъекции нет.
func patchUser(db *sql.DB, id int, name *string, age *int) error {
	parts := []string{} // куски вида "name = $1"
	args := []any{}
	n := 1 // номер следующего плейсхолдера

	if name != nil {
		parts = append(parts, fmt.Sprintf("name = $%d", n))
		args = append(args, *name)
		n++
	}
	if age != nil {
		parts = append(parts, fmt.Sprintf("age = $%d", n))
		args = append(args, *age)
		n++
	}
	if len(parts) == 0 {
		return fmt.Errorf("нечего обновлять")
	}

	// получится, например: UPDATE users SET name = $1, age = $2 WHERE id = $3
	query := "UPDATE users SET " + strings.Join(parts, ", ") + fmt.Sprintf(" WHERE id = $%d", n)
	args = append(args, id)

	fmt.Println("SQL:", query, args)
	_, err := db.Exec(query, args...)
	return err
}

func main() {
	db := mustConnect()
	defer db.Close()

	fmt.Println("--- Обновляем возраст Али ---")
	if err := updateUserAge(db, 1, 26); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Обновляем несуществующего пользователя ---")
	if err := updateUserAge(db, 999, 30); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Обновляем несколько полей Веры ---")
	err := updateUser(db, User{ID: 2, Name: "Вера", Email: "vera@mail.tj", Age: 31})
	if err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Обновлено")
	}

	fmt.Println("\n--- Частичное обновление Тимура ---")
	name := "Тимур"
	age := 20
	if err := patchUser(db, 3, &name, &age); err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Обновлены имя и возраст")
	}
	if err := patchUser(db, 3, nil, &age); err != nil { // только возраст
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Обновлён только возраст")
	}
	if err := patchUser(db, 3, nil, nil); err != nil { // ничего не задано
		fmt.Println("✗", err)
	}
}
