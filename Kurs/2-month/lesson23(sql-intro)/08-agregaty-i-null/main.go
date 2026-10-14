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

// getStats считает статистику по пользователям не младше minAge.
// Агрегаты (COUNT, AVG, MIN, MAX) возвращают ровно одну строку - значит, QueryRow.
func getStats(db *sql.DB, minAge int) error {
	var count int

	// NULL - это не 0 и не пустая строка, а "значения нет".
	// Обычный float64 не умеет хранить NULL, поэтому берём sql.NullFloat64:
	// у него есть поле Valid ("значение есть?") и поле Float64 (само значение)
	var avgAge, lowest, highest sql.NullFloat64

	err := db.QueryRow(`
		SELECT COUNT(*), AVG(age), MIN(age), MAX(age)
		FROM users
		WHERE age >= $1
	`, minAge).Scan(&count, &avgAge, &lowest, &highest)
	if err != nil {
		return err
	}

	fmt.Println("Всего пользователей:", count)

	if avgAge.Valid {
		fmt.Printf("Средний возраст: %.1f\n", avgAge.Float64)
		fmt.Printf("Диапазон: %.0f - %.0f\n", lowest.Float64, highest.Float64)
	} else {
		// Если подходящих строк нет, AVG/MIN/MAX возвращают NULL
		fmt.Println("Нет данных о возрасте")
	}
	return nil
}

func main() {
	db := mustConnect()
	defer db.Close()

	fmt.Println("=== Все пользователи ===")
	if err := getStats(db, 0); err != nil {
		log.Fatal(err)
	}

	// Никому нет 1000 лет: подходящих строк нет, и AVG вернёт NULL
	fmt.Println("\n=== Пользователи от 1000 лет ===")
	if err := getStats(db, 1000); err != nil {
		log.Fatal(err)
	}

	// Для NULL-колонок есть: sql.NullString, sql.NullInt64, sql.NullFloat64, sql.NullBool.
	// Вместо них можно использовать указатели: *string, *int (nil = NULL)
}
