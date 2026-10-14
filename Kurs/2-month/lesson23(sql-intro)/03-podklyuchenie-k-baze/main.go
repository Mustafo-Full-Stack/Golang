package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Строка подключения - набор пар ключ=значение:
	//   host     - где работает база (localhost = этот же компьютер)
	//   port     - порт PostgreSQL (по умолчанию 5432)
	//   user     - имя пользователя
	//   password - пароль
	//   dbname   - имя базы
	//   sslmode  - disable = без шифрования (для локальной разработки это нормально)
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=coursedb sslmode=disable"

	// sql.Open НЕ подключается к базе! Он только создаёт пул соединений.
	// "postgres" - имя драйвера, который зарегистрировал pq
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("не удалось создать пул:", err)
	}
	defer db.Close()

	// Ping - настоящая проверка: достучались ли мы до базы
	if err := db.Ping(); err != nil {
		log.Fatal("не удалось подключиться:", err)
	}

	fmt.Println("Подключение к базе установлено!")

	// Эксперименты:
	// 1) поменяйте password=secret на неверный пароль - Ping вернёт понятную ошибку
	// 2) поменяйте dbname на несуществующую базу
	// 3) остановите PostgreSQL и запустите программу
}
