package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Мини-проект: интернет-магазин.
// Покупка - это сразу несколько действий: создать заказ, списать товар со склада,
// списать деньги со счёта покупателя. Любая ошибка должна откатить ВСЁ.

type OrderItem struct {
	ProductID int
	Quantity  int
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

func withTransaction(db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("ошибка %v, откат тоже не удался: %w", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// placeOrder оформляет заказ: пять операций в одной транзакции
func placeOrder(db *sql.DB, accountID int, items []OrderItem) (int, error) {
	var orderID int

	err := withTransaction(db, func(tx *sql.Tx) error {
		// 1. Создаём заказ
		err := tx.QueryRow(`
			INSERT INTO orders (account_id, status) VALUES ($1, 'pending') RETURNING id
		`, accountID).Scan(&orderID)
		if err != nil {
			return fmt.Errorf("создание заказа: %w", err)
		}

		total := 0.0

		// 2. Обрабатываем каждый товар
		for _, item := range items {
			var name string
			var price float64
			var stock int

			// Блокируем строку товара, чтобы никто другой не купил его одновременно с нами
			err := tx.QueryRow(`
				SELECT name, price, quantity FROM products WHERE id = $1 FOR UPDATE
			`, item.ProductID).Scan(&name, &price, &stock)
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("товар %d не найден", item.ProductID)
			}
			if err != nil {
				return err
			}

			if stock < item.Quantity {
				return fmt.Errorf("недостаточно товара %q: на складе %d, нужно %d",
					name, stock, item.Quantity)
			}

			// Уменьшаем остаток на складе
			_, err = tx.Exec(`
				UPDATE products SET quantity = quantity - $1 WHERE id = $2
			`, item.Quantity, item.ProductID)
			if err != nil {
				return fmt.Errorf("списание остатка: %w", err)
			}

			// Добавляем позицию заказа
			_, err = tx.Exec(`
				INSERT INTO order_items (order_id, product_id, quantity, price)
				VALUES ($1, $2, $3, $4)
			`, orderID, item.ProductID, item.Quantity, price)
			if err != nil {
				return fmt.Errorf("добавление позиции: %w", err)
			}

			total += price * float64(item.Quantity)
		}

		// 3. Списываем деньги со счёта покупателя
		var balance float64
		err = tx.QueryRow(`
			SELECT balance FROM accounts WHERE id = $1 FOR UPDATE
		`, accountID).Scan(&balance)
		if err != nil {
			return fmt.Errorf("чтение счёта: %w", err)
		}
		if balance < total {
			return fmt.Errorf("недостаточно средств: на счёте %.2f, нужно %.2f", balance, total)
		}

		_, err = tx.Exec(`
			UPDATE accounts SET balance = balance - $1 WHERE id = $2
		`, total, accountID)
		if err != nil {
			return fmt.Errorf("списание денег: %w", err)
		}

		// 4. Финализируем заказ
		_, err = tx.Exec(`
			UPDATE orders SET status = 'paid', total = $1 WHERE id = $2
		`, total, orderID)
		if err != nil {
			return fmt.Errorf("обновление заказа: %w", err)
		}

		fmt.Printf("Заказ #%d оформлен на сумму %.2f\n", orderID, total)
		return nil
	})

	// Если была ошибка, заказ откатился - его номера не существует
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

// printState показывает всё, что может измениться при покупке
func printState(db *sql.DB) {
	var balance float64
	err := db.QueryRow(`SELECT balance FROM accounts WHERE id = 1`).Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	var ordersCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM orders`).Scan(&ordersCount)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT name, quantity FROM products ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("  Счёт Али: %.2f | заказов в базе: %d\n  Остатки:", balance, ordersCount)
	for rows.Next() {
		var name string
		var quantity int
		if err := rows.Scan(&name, &quantity); err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" %s=%d", name, quantity)
	}
	fmt.Println()
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := mustConnect()
	defer db.Close()

	// Товары по id: 1 - Ноутбук (800), 2 - Мышь (25.50), 3 - Клавиатура (45), 4 - Монитор (300)
	fmt.Println("Начальное состояние:")
	printState(db)

	fmt.Println("\n--- 1. Успешная покупка: ноутбук + 2 мыши ---")
	_, err := placeOrder(db, 1, []OrderItem{{ProductID: 1, Quantity: 1}, {ProductID: 2, Quantity: 2}})
	if err != nil {
		fmt.Println("✗", err)
	}
	printState(db)

	fmt.Println("\n--- 2. Не хватает товара: 10 клавиатур (на складе 3) ---")
	_, err = placeOrder(db, 1, []OrderItem{{ProductID: 3, Quantity: 10}})
	if err != nil {
		fmt.Println("✗", err)
	}
	printState(db) // ничего не изменилось

	fmt.Println("\n--- 3. Не хватает денег: 2 ноутбука + 2 монитора ---")
	// Остатки к этому моменту уже уменьшены внутри транзакции,
	// но деньги не списать - и откат возвращает всё на склад
	_, err = placeOrder(db, 1, []OrderItem{{ProductID: 1, Quantity: 2}, {ProductID: 4, Quantity: 2}})
	if err != nil {
		fmt.Println("✗", err)
	}
	printState(db) // остатки и деньги прежние

	fmt.Println("\n--- 4. Несуществующий товар ---")
	_, err = placeOrder(db, 1, []OrderItem{{ProductID: 999, Quantity: 1}})
	if err != nil {
		fmt.Println("✗", err)
	}
	printState(db)
}
