package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Шаг 6. Всё, что идёт по сети, может сломаться. Три уровня проверки:
//  1. err от http.Get  — нет интернета, неверный адрес, таймаут
//  2. resp.StatusCode  — сервер ответил, но "не 200" (например, 404)
//  3. err от Decode    — пришёл не тот JSON
func getUser(id int) (User, error) {
	// Таймаут: не ждём ответ вечно
	client := http.Client{Timeout: 5 * time.Second}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	resp, err := client.Get(url)
	if err != nil {
		return User{}, fmt.Errorf("запрос: %w", err)
	}
	defer resp.Body.Close()

	// 404 — это НЕ ошибка Go. Запрос прошёл успешно, просто там пусто.
	// Если не проверить статус, мы попытаемся разобрать "{}" и получим пустого юзера.
	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("сервер вернул статус %s", resp.Status)
	}

	var u User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return User{}, fmt.Errorf("разбор JSON: %w", err)
	}
	return u, nil
}

func main() {
	// id=1 существует, id=999 — нет (будет 404)
	for _, id := range []int{1, 999} {
		u, err := getUser(id)
		if err != nil {
			fmt.Printf("id=%d: ошибка: %v\n", id, err)
			continue
		}
		fmt.Printf("id=%d: найден %s\n", id, u.Name)
	}
}
