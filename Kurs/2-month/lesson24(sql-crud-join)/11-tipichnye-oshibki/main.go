package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
)

// Типичные ошибки при изменении данных.
// Код здесь нарочно "неправильный" - мы смотрим, как ошибка проявляется.
// Не исправляйте его, ошибка и есть урок.
//
// Для ошибки 4 нужна таблица orders (создаётся на шаге 7).

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

// mustExec - выполнить запрос или завершить программу (чтобы не писать проверку каждый раз)
func mustExec(db *sql.DB, query string) {
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Ошибка 1: LastInsertId в PostgreSQL
func mistakeLastInsertID(db *sql.DB) {
	fmt.Println("--- 1. LastInsertId в PostgreSQL ---")

	// Запрос, который ничего не меняет: просто чтобы получить sql.Result
	res, err := db.Exec(`UPDATE users SET age = age WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	fmt.Println("id:", id, "| ошибка:", err)
	// Решение: INSERT ... RETURNING id и QueryRow().Scan(&id)
}

// Ошибка 2: не проверили RowsAffected
func mistakeNoRowsAffected(db *sql.DB) {
	fmt.Println("\n--- 2. Не проверили RowsAffected ---")

	// Пользователя 99999 не существует, но ошибки нет
	res, err := db.Exec(`UPDATE users SET age = 30 WHERE id = 99999`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Обновлено!") // ...но на самом деле ничего не обновилось

	n, _ := res.RowsAffected()
	fmt.Println("На самом деле изменено строк:", n)
}

// Ошибка 3: UPDATE без WHERE - изменились ВСЕ строки.
// Опыт на отдельной тестовой таблице, чтобы не испортить users.
func mistakeNoWhere(db *sql.DB) {
	fmt.Println("\n--- 3. UPDATE без WHERE ---")

	mustExec(db, `DROP TABLE IF EXISTS demo_prices`)
	mustExec(db, `CREATE TABLE demo_prices (id SERIAL PRIMARY KEY, item TEXT, price INT)`)
	mustExec(db, `INSERT INTO demo_prices (item, price) VALUES ('Хлеб', 5), ('Молоко', 12), ('Сыр', 40)`)

	// Хотели обнулить цену только у сыра, но забыли WHERE id = 3
	res, err := db.Exec(`UPDATE demo_prices SET price = 0`)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("Изменено строк:", n, "(а хотели 1)")

	printPrices(db) // все цены стали 0

	mustExec(db, `DROP TABLE demo_prices`) // убираем за собой
}

func printPrices(db *sql.DB) {
	rows, err := db.Query(`SELECT item, price FROM demo_prices ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item string
		var price int
		if err := rows.Scan(&item, &price); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  %-8s %d\n", item, price)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// Ошибка 4: нарушение внешнего ключа
func mistakeForeignKey(db *sql.DB) {
	fmt.Println("\n--- 4. Нарушение внешнего ключа ---")

	// Пользователя 999 не существует
	_, err := db.Exec(`INSERT INTO orders (user_id, product, amount) VALUES (999, 'Товар', 10)`)
	fmt.Println("Ошибка:", err)

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		fmt.Println("Код:", pqErr.Code, "-", pqErr.Code.Name())
	}
}

// Ошибка 5: Query вместо Exec для INSERT/UPDATE/DELETE
func mistakeQueryInsteadOfExec(db *sql.DB) {
	fmt.Println("\n--- 5. Query вместо Exec ---")

	// Query возвращает rows, и мы обязаны их закрыть. Мы не закрыли - соединение висит
	rows, err := db.Query(`UPDATE users SET age = age WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	_ = rows // забыли rows.Close()

	fmt.Println("Соединений занято:", db.Stats().InUse) // 1 - утечка
	// Решение: для запросов без выборки всегда db.Exec
}

// Ещё одна ошибка, которую нельзя показать в коде:
//
// 6. Забыли stmt.Close() у подготовленного выражения (db.Prepare).
//    Утечка ресурсов на стороне сервера БД. Решение: defer stmt.Close() сразу после Prepare.

func main() {
	db := mustConnect()
	defer db.Close()

	mistakeLastInsertID(db)
	mistakeNoRowsAffected(db)
	mistakeNoWhere(db)
	mistakeForeignKey(db)
	mistakeQueryInsteadOfExec(db) // последней: она занимает соединение
}
