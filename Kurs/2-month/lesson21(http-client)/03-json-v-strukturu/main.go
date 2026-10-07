package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Шаг 3. Ответ сервера — JSON. Разбираем его прямо в структуру.
func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Сервер вернул:", resp.Status)
		return
	}

	var todo Todo
	// NewDecoder читает прямо из resp.Body, промежуточный []byte не нужен
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		fmt.Println("Ошибка разбора:", err)
		return
	}

	fmt.Printf("наша струкртура todo:%+v\n", todo)
	fmt.Printf("Задача #%d: %s\n", todo.ID, todo.Title)
	fmt.Printf("Выполнена: %t\n", todo.Completed)
}
