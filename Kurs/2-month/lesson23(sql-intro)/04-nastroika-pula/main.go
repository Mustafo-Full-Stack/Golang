package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=coursedb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("не удалось создать пул:", err)
	}
	defer db.Close()

	// Настройка пула соединений.
	// Без неё пул может открыть слишком много соединений, и база откажет.
	db.SetMaxOpenConns(25)                 // максимум соединений одновременно
	db.SetMaxIdleConns(5)                  // сколько держать "про запас" (простаивающими)
	db.SetConnMaxLifetime(5 * time.Minute) // сколько живёт одно соединение

	// Статистика пула - можно посмотреть, что происходит внутри.
	// Сначала соединений нет, потому что Open ничего не открывает:
	fmt.Println("До Ping:    ", db.Stats().OpenConnections, "соединений")

	if err := db.Ping(); err != nil {
		log.Fatal("не удалось подключиться:", err)
	}

	// А после Ping появилось одно соединение, и оно простаивает в пуле
	stats := db.Stats()
	fmt.Println("После Ping:  ", stats.OpenConnections, "соединений")
	fmt.Println("Из них занято:", stats.InUse)
	fmt.Println("Простаивает:  ", stats.Idle)
	fmt.Println("Максимум:     ", stats.MaxOpenConnections)
}
