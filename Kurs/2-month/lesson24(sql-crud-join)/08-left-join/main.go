package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UserStats struct {
	Name        string
	OrdersCount int
	TotalSpent  sql.NullFloat64 // у тех, кто ничего не заказывал, SUM будет NULL
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

// LEFT JOIN берёт ВСЕХ пользователей из левой таблицы (users),
// а заказы подклеивает, если они есть. Нет заказов - вместо них NULL.
// (INNER JOIN такого пользователя просто не показал бы.)
func getUserStats(db *sql.DB) ([]UserStats, error) {
	rows, err := db.Query(`
		SELECT u.name, COUNT(o.id) AS orders_count, SUM(o.amount) AS total
		FROM users u
		LEFT JOIN orders o ON u.id = o.user_id
		GROUP BY u.id, u.name
		ORDER BY total DESC NULLS LAST
	`)
	if err != nil {
		return nil, fmt.Errorf("запрос статистики: %w", err)
	}
	defer rows.Close()

	var stats []UserStats
	for rows.Next() {
		var s UserStats
		if err := rows.Scan(&s.Name, &s.OrdersCount, &s.TotalSpent); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return stats, rows.Err()
}

func main() {
	db := mustConnect()
	defer db.Close()

	stats, err := getUserStats(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range stats {
		total := 0.0
		if s.TotalSpent.Valid {
			total = s.TotalSpent.Float64
		}
		fmt.Printf("%-10s заказов: %d, потрачено: %.2f\n", s.Name, s.OrdersCount, total)
	}
}
