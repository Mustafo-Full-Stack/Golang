package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
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

// Вариант 1: ГОНКА. Читаем баланс, считаем в Go, записываем результат.
// Между чтением и записью другие горутины успевают сделать то же самое.
func withdrawBroken(db *sql.DB) {
	var balance float64
	err := db.QueryRow(`SELECT balance FROM accounts WHERE id = 1`).Scan(&balance)
	if err != nil {
		log.Println("чтение:", err)
		return
	}

	time.Sleep(10 * time.Millisecond) // "думаем" - в это время другие тоже читают то же число

	_, err = db.Exec(`UPDATE accounts SET balance = $1 WHERE id = 1`, balance-100)
	if err != nil {
		log.Println("запись:", err)
	}
}

// Вариант 2: атомарное обновление прямо в SQL - вычитает сама база
func withdrawAtomic(db *sql.DB) {
	_, err := db.Exec(`UPDATE accounts SET balance = balance - 100 WHERE id = 1`)
	if err != nil {
		log.Println("запись:", err)
	}
}

// Вариант 3: транзакция + SELECT ... FOR UPDATE.
// Строка заблокирована: остальные горутины ждут своей очереди.
func withdrawLocked(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("начало транзакции:", err)
		return
	}
	defer tx.Rollback()

	var balance float64
	err = tx.QueryRow(`SELECT balance FROM accounts WHERE id = 1 FOR UPDATE`).Scan(&balance)
	if err != nil {
		log.Println("чтение:", err)
		return
	}

	time.Sleep(10 * time.Millisecond) // то же самое "думание"

	_, err = tx.Exec(`UPDATE accounts SET balance = $1 WHERE id = 1`, balance-100)
	if err != nil {
		log.Println("запись:", err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println("commit:", err)
	}
}

// runParallel сбрасывает баланс Али до 1000, запускает work в 10 горутинах
// и показывает итог. 10 списаний по 100 должны дать ровно 0.
func runParallel(db *sql.DB, name string, work func()) {
	_, err := db.Exec(`UPDATE accounts SET balance = 1000 WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for n := 1; n <= 10; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()

	var final float64
	err = db.QueryRow(`SELECT balance FROM accounts WHERE id = 1`).Scan(&final)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-22s ожидали 0.00, получили %.2f\n", name, final)
}

func main() {
	db := mustConnect()
	defer db.Close()

	// 10 горутин одновременно списывают с одного счёта по 100: 1000 - 10*100 = 0
	runParallel(db, "Без блокировки:", func() { withdrawBroken(db) })
	runParallel(db, "Атомарный UPDATE:", func() { withdrawAtomic(db) })
	runParallel(db, "Транзакция FOR UPDATE:", func() { withdrawLocked(db) })

	// Первый вариант - классическая гонка (как в уроке про горутины и мьютексы):
	// почти все списания "потерялись". База данных - такой же общий ресурс,
	// и её нужно защищать: атомарным запросом или блокировкой строки.
	//
	// В конце вернём баланс как было
	db.Exec(`UPDATE accounts SET balance = 1000 WHERE id = 1`)
}
