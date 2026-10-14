package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Заказ вместе с данными клиента - результат склейки двух таблиц
type OrderWithUser struct {
	OrderID   int
	Product   string
	Amount    float64
	UserName  string
	UserEmail string
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

// JOIN - склейка двух таблиц по общему признаку.
// В заказе хранится только user_id (число), а мы хотим увидеть имя клиента.
// o и u - короткие псевдонимы таблиц (orders и users).
func getOrdersWithUsers(db *sql.DB) ([]OrderWithUser, error) {
	rows, err := db.Query(`
		SELECT o.id, o.product, o.amount, u.name, u.email
		FROM orders o
		INNER JOIN users u ON o.user_id = u.id
		ORDER BY o.id
	`)
	if err != nil {
		return nil, fmt.Errorf("запрос заказов: %w", err)
	}
	defer rows.Close()

	var result []OrderWithUser
	for rows.Next() {
		var o OrderWithUser
		err := rows.Scan(&o.OrderID, &o.Product, &o.Amount, &o.UserName, &o.UserEmail)
		if err != nil {
			return nil, err
		}
		result = append(result, o)
	}

	return result, rows.Err()
}

func main() {
	db := mustConnect()
	defer db.Close()

	orders, err := getOrdersWithUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-4s %-14s %10s  %-10s\n", "ID", "Товар", "Сумма", "Клиент")
	for _, o := range orders {
		fmt.Printf("%-4d %-14s %10.2f  %-10s\n", o.OrderID, o.Product, o.Amount, o.UserName)
	}
}
