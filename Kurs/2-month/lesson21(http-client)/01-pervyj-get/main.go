package main

import (
	"fmt"
	"io"
	"net/http"
)

// Шаг 1. Первый GET-запрос: наша программа сама ходит в интернет.
// Браузер делает то же самое, когда вы открываете сайт.
func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close() // СРАЗУ после проверки err, пока не забыли

	// Body — это поток данных, читаем всё в []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Println("Статус:", resp.Status)
	fmt.Println("Статус код:", resp.StatusCode)
	fmt.Println("Тело:", string(body))

	// Почему defer именно после проверки err?
	// Если запрос не удался, resp == nil, и resp.Body.Close() вызовет панику.
}
