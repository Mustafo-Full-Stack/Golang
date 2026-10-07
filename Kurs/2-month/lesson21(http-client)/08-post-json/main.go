package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

// Что отправляем
type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// Что сервер возвращает: то же самое, но уже с номером (ID)
type CreatedPost struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// Шаг 8. POST-запрос: отправляем JSON, получаем JSON обратно.
func createPost(p NewPost) (*CreatedPost, error) {
	// 1. Структура -> JSON ([]byte)
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("сериализация: %w", err)
	}

	// 2. Отправляем. bytes.NewBuffer превращает []byte в io.Reader,
	//    а именно его ждёт метод Post. "application/json" — это заголовок Content-Type.
	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("запрос: %w", err)
	}
	defer resp.Body.Close()

	// 3. При создании сервер отвечает 201 Created (а не 200)
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("ожидали 201, получили %d", resp.StatusCode)
	}

	// 4. JSON из ответа -> структура
	var created CreatedPost
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, fmt.Errorf("разбор: %w", err)
	}
	return &created, nil
}

func main() {
	p, err := createPost(NewPost{
		Title:  "Изучаю Go",
		Body:   "Сегодня научился делать HTTP-запросы",
		UserID: 1,
	})
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Создан пост #%d: %s\n", p.ID, p.Title)
}
