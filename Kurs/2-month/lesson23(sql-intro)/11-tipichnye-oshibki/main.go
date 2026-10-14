package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Типичные ошибки при работе с database/sql.
// Код здесь нарочно "неправильный" - мы смотрим, как ошибка проявляется.
// Не исправляйте его, ошибка и есть урок.

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

// Ошибка 1: порядок в Scan не совпадает с порядком в SELECT
func mistakeScanOrder(db *sql.DB) {
	fmt.Println("--- 1. Порядок в Scan ---")

	var id, age int
	var name string

	// В SELECT: id, name, age. А в Scan перепутали: name, id, age
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = 1").
		Scan(&name, &id, &age)

	fmt.Println("Ошибка:", err)
	// Строка "Али" не превращается в число - Scan вернул ошибку
}

// Ошибка 2: NULL в обычный тип
func mistakeNullIntoFloat(db *sql.DB) {
	fmt.Println("\n--- 2. NULL в обычный float64 ---")

	// Нет ни одного пользователя старше 1000 лет, поэтому AVG вернёт NULL
	var avg float64
	err := db.QueryRow("SELECT AVG(age) FROM users WHERE age > 1000").Scan(&avg)

	fmt.Println("Ошибка:", err)
	// Решение: var avg sql.NullFloat64 (и проверять avg.Valid)
}

// Ошибка 3: забыли rows.Close() - соединения утекают
func mistakeForgotClose(db *sql.DB) {
	fmt.Println("\n--- 3. Забыли rows.Close() ---")

	// Сначала правильно: 5 запросов, и каждый закрываем
	for n := 1; n <= 5; n++ {
		rows, err := db.Query("SELECT id FROM users")
		if err != nil {
			log.Fatal(err)
		}
		rows.Next() // прочитали одну строку
		rows.Close()
	}
	fmt.Println("С Close: соединений занято =", db.Stats().InUse) // 0

	// Теперь те же 5 запросов, но Close забыли
	for n := 1; n <= 5; n++ {
		rows, err := db.Query("SELECT id FROM users")
		if err != nil {
			log.Fatal(err)
		}
		rows.Next() // прочитали одну строку, а Close нет
	}
	fmt.Println("Без Close: соединений занято =", db.Stats().InUse) // 5

	// Каждое забытое соединение навсегда выпало из пула.
	// В цикле на 1000 запросов соединения кончатся, и программа зависнет.
}

// Другие ошибки (их нельзя показать в одной программе, показываем отдельно):
//
// 4. Нет "_" перед драйвером.
//    import "github.com/lib/pq"   -> ошибка компиляции: imported and not used
//    Убрали импорт совсем         -> sql: unknown driver "postgres" (forgotten import?)
//
// 5. Не проверили rows.Err() после цикла.
//    Ошибка посреди обхода останется незамеченной, вернутся неполные данные.
//
// 6. sql.Open в каждой функции.
//    Получается куча пулов. Нужен один *sql.DB на всё приложение.

func main() {
	db := mustConnect()
	defer db.Close()

	mistakeScanOrder(db)
	mistakeNullIntoFloat(db)
	mistakeForgotClose(db)
}
