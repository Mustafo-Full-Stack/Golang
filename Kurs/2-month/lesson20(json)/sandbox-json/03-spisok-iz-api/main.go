package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"` // в JSON именно userId (camelCase)
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Шаг 3. API вернул МАССИВ -> разбираем в срез.
// Новинка: json.NewDecoder читает прямо из resp.Body, без io.ReadAll.
func main() {
	// ?_limit=5 — просим сервер вернуть только 5 задач
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos?_limit=5")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	for _, t := range todos {
		mark := "[ ]"
		if t.Completed {
			mark = "[x]"
		}
		fmt.Println(mark, t.ID, t.Title)
	}

	// Resp.Body — это io.Reader, Decoder умеет читать из любого io.Reader
	// (файл, сеть, строка). Поэтому код получился короче.
}
