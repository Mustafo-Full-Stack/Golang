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

// deleteUser удаляет строку навсегда
func deleteUser(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("удаление: %w", err)
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("пользователь с id=%d не найден", id)
	}
	return nil
}

// softDeleteUser - "мягкое" удаление: строка остаётся, мы только помечаем её неактивной.
// В реальных системах данные редко удаляют физически: так можно восстановить
// и сохранить историю.
func softDeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec(`UPDATE users SET is_active = false WHERE id = $1`, id)
	return err
}

func restoreUser(db *sql.DB, id int) error {
	_, err := db.Exec(`UPDATE users SET is_active = true WHERE id = $1`, id)
	return err
}

func isActive(db *sql.DB, id int) bool {
	var active bool
	err := db.QueryRow(`SELECT is_active FROM users WHERE id = $1`, id).Scan(&active)
	if err != nil {
		log.Fatal(err)
	}
	return active
}

func main() {
	db := mustConnect()
	defer db.Close()

	// Создаём временного пользователя, чтобы было что удалять
	id, err := createUserReturningID(db, "Временный", "temp@mail.tj", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создан пользователь с ID:", id)

	fmt.Println("\n--- Удаляем его ---")
	if err := deleteUser(db, id); err != nil {
		fmt.Println("✗", err)
	} else {
		fmt.Println("✓ Удалён")
	}

	fmt.Println("\n--- Удаляем ещё раз ---")
	if err := deleteUser(db, id); err != nil {
		fmt.Println("✗", err)
	}

	fmt.Println("\n--- Мягкое удаление Нигины (id=4) ---")
	softDeleteUser(db, 4)
	fmt.Println("Активна:", isActive(db, 4)) // false - но строка на месте

	restoreUser(db, 4) // и её можно восстановить
	fmt.Println("После восстановления активна:", isActive(db, 4))

	// СТРАШНЫЙ ПРИМЕР - не запускайте!
	// Забытый WHERE в DELETE или UPDATE уничтожает данные всей таблицы за секунду:
	//
	//     db.Exec("DELETE FROM users")
	//
	// Всегда сначала проверяйте условие запросом SELECT.
}
