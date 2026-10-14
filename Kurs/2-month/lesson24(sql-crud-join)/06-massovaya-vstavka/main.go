package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

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

// Способ 1: обычный цикл - на каждую вставку отдельный запрос
func insertManySlow(db *sql.DB, names []string) error {
	start := time.Now()
	for _, n := range names {
		_, err := db.Exec(`INSERT INTO tags (name) VALUES ($1)`, n)
		if err != nil {
			return err
		}
	}
	fmt.Println("Обычная вставка:", time.Since(start))
	return nil
}

// Способ 2: подготовленное выражение (prepared statement).
// Запрос "подготавливаем" один раз, потом только подставляем значения.
// База разбирает и планирует запрос единожды.
func insertManyPrepared(db *sql.DB, names []string) error {
	stmt, err := db.Prepare(`INSERT INTO tags (name) VALUES ($1)`)
	if err != nil {
		return err
	}
	defer stmt.Close() // подготовленное выражение обязательно закрываем

	start := time.Now()
	for _, n := range names {
		if _, err := stmt.Exec(n); err != nil {
			return err
		}
	}
	fmt.Println("Prepared:       ", time.Since(start))
	return nil
}

// Способ 3: один запрос на все строки сразу:
// INSERT INTO tags (name) VALUES ($1), ($2), ($3), ...
func insertBatch(db *sql.DB, names []string) error {
	if len(names) == 0 {
		return nil
	}

	start := time.Now()
	values := make([]string, 0, len(names))
	args := make([]any, 0, len(names))
	for i, n := range names {
		values = append(values, fmt.Sprintf("($%d)", i+1))
		args = append(args, n)
	}

	query := "INSERT INTO tags (name) VALUES " + strings.Join(values, ",")
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	fmt.Println("Один запрос:    ", time.Since(start))
	return nil
}

// TRUNCATE - быстро очищает таблицу целиком (это DDL-команда: меняет таблицу, а не строки)
func clearTags(db *sql.DB) {
	if _, err := db.Exec(`TRUNCATE TABLE tags`); err != nil {
		log.Fatal(err)
	}
}

func countTags(db *sql.DB) int {
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM tags`).Scan(&count); err != nil {
		log.Fatal(err)
	}
	return count
}

func main() {
	db := mustConnect()
	defer db.Close()

	// Готовим 1000 имён: tag-1, tag-2, ...
	names := []string{}
	for n := 1; n <= 1000; n++ {
		names = append(names, fmt.Sprintf("tag-%d", n))
	}

	clearTags(db)
	if err := insertManySlow(db, names); err != nil {
		log.Fatal(err)
	}
	fmt.Println("  строк в таблице:", countTags(db))

	clearTags(db)
	if err := insertManyPrepared(db, names); err != nil {
		log.Fatal(err)
	}
	fmt.Println("  строк в таблице:", countTags(db))

	clearTags(db)
	if err := insertBatch(db, names); err != nil {
		log.Fatal(err)
	}
	fmt.Println("  строк в таблице:", countTags(db))

	// Результат одинаковый (1000 строк), а время сильно разное.
}
