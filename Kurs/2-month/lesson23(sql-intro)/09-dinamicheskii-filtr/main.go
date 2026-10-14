package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Age       int
	IsActive  bool
	CreatedAt time.Time
}

// Filter - какие условия задал пользователь.
// Указатель нужен, чтобы отличить "не задано" (nil) от "задано нулевое значение" (0, false).
// Это те же указатели, что мы проходили раньше: *int = "фильтр не задан"
type Filter struct {
	MinAge   *int
	MaxAge   *int
	IsActive *bool
	NameLike *string
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

func searchUsers(db *sql.DB, f Filter) ([]User, error) {
	// WHERE 1=1 - всегда истинное условие.
	// Благодаря ему к запросу можно дописывать " AND ..." не думая, первое это условие или нет
	query := `SELECT id, name, email, age, is_active, created_at FROM users WHERE 1=1`
	args := []any{}
	argNum := 1 // номер следующего плейсхолдера: $1, $2, $3...

	if f.MinAge != nil {
		query += fmt.Sprintf(" AND age >= $%d", argNum)
		args = append(args, *f.MinAge)
		argNum++
	}
	if f.MaxAge != nil {
		query += fmt.Sprintf(" AND age <= $%d", argNum)
		args = append(args, *f.MaxAge)
		argNum++
	}
	if f.IsActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argNum)
		args = append(args, *f.IsActive)
		argNum++
	}
	if f.NameLike != nil {
		// ILIKE - как LIKE, но без учёта регистра
		query += fmt.Sprintf(" AND name ILIKE $%d", argNum)
		args = append(args, "%"+*f.NameLike+"%")
		argNum++
	}
	query += " ORDER BY id"

	// Смотрим, какой запрос получился
	fmt.Println("SQL:", query)
	fmt.Println("Параметры:", args)

	// В текст запроса подставляются ТОЛЬКО номера плейсхолдеров ($1, $2...).
	// Сами значения всегда идут через args - значит, SQL-инъекция по-прежнему невозможна.
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func printNames(users []User) {
	for _, u := range users {
		fmt.Printf("  %s, %d лет\n", u.Name, u.Age)
	}
	fmt.Println()
}

func main() {
	db := mustConnect()
	defer db.Close()

	// 1. Фильтр пустой - вернутся все
	fmt.Println("=== Без фильтров ===")
	users, err := searchUsers(db, Filter{})
	if err != nil {
		log.Fatal(err)
	}
	printNames(users)

	// 2. Только минимальный возраст
	minAge := 20
	fmt.Println("=== Возраст от 20 ===")
	users, err = searchUsers(db, Filter{MinAge: &minAge})
	if err != nil {
		log.Fatal(err)
	}
	printNames(users)

	// 3. Диапазон возраста + активные
	maxAge := 28
	active := true
	fmt.Println("=== Возраст от 20 до 28, только активные ===")
	users, err = searchUsers(db, Filter{MinAge: &minAge, MaxAge: &maxAge, IsActive: &active})
	if err != nil {
		log.Fatal(err)
	}
	printNames(users)

	// 4. Поиск по части имени (регистр не важен)
	part := "ВЕР"
	fmt.Println("=== Имя содержит ВЕР ===")
	users, err = searchUsers(db, Filter{NameLike: &part})
	if err != nil {
		log.Fatal(err)
	}
	printNames(users)
}
