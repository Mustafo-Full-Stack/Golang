package main

import (
	"fmt"
	"net/http"
	"time"
)

// Шаг 5. Свой клиент с таймаутом.
// http.Get использует клиент БЕЗ таймаута: сервер завис — программа ждёт вечно.
// В боевом коде всегда создаём свой клиент.
var client = &http.Client{
	Timeout: 5 * time.Second,
}

func main() {
	// Нормальный запрос: успеваем за 5 секунд
	resp, err := client.Get("https://api.github.com/users/golang")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Статус:", resp.Status)

	// А теперь специально даём всего 1 миллисекунду — не успеем
	fast := &http.Client{Timeout: 1 * time.Millisecond}

	resp2, err := fast.Get("https://api.github.com/users/golang")
	if err != nil {
		fmt.Println("Ошибка (таймаут):", err)
		return
	}
	defer resp2.Body.Close()
	fmt.Println("Статус:", resp2.Status)
}
