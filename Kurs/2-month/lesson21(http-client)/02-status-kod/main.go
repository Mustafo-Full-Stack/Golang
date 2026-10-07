package main

import (
	"fmt"
	"net/http"
)

// Шаг 2. Статус-код: сервер ответил, но всё ли в порядке?
// 2xx — успех, 4xx — виноват ты, 5xx — виноват сервер.
func check(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%s -> %s\n", url, resp.Status)

	switch {
	case resp.StatusCode == http.StatusOK:
		fmt.Println("  Всё хорошо")
	case resp.StatusCode == http.StatusNotFound:
		fmt.Println("  Ресурс не найден")
	case resp.StatusCode == http.StatusUnauthorized:
		fmt.Println("  Нужна авторизация")
	case resp.StatusCode >= 500:
		fmt.Println("  Проблема на сервере, попробуйте позже")
	case resp.StatusCode >= 400:
		fmt.Println("  Ошибка в запросе")
	}
}

func main() {
	check("https://jsonplaceholder.typicode.com/todos/1")     // 200
	check("https://jsonplaceholder.typicode.com/todos/99999") // 404
	check("https://httpbin.org/status/503")                   // 503

	// Важно: 404 — это НЕ ошибка Go. err == nil, запрос прошёл успешно,
	// просто сервер ответил "такого нет". Поэтому статус проверяем сами.
}
