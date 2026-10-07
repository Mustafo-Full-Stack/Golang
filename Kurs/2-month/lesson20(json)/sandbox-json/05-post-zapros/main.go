package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// То, что МЫ отправляем на сервер
type NewPost struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// То, что сервер присылает в ответ (с присвоенным ID)
type PostResponse struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Шаг 5. Теперь в обе стороны:
// Go-структура -> Marshal -> отправляем -> получаем JSON -> Unmarshal -> Go-структура
func main() {
	post := NewPost{UserID: 1, Title: "Привет из Go", Body: "Я отправил это из своей программы"}

	// 1. Упаковали структуру в JSON
	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	// 2. Отправили POST (отдать данные серверу).
	// bytes.NewReader превращает []byte в io.Reader — то, что нужно для отправки.
	// "application/json" говорит серверу: "я шлю JSON"
	resp, err := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Статус:", resp.Status) // 201 Created — "создано"

	// 3. Разобрали ответ сервера
	var created PostResponse
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	fmt.Printf("Сервер создал пост с ID=%d: %q\n", created.ID, created.Title)
}
