package main

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

// Шаг 6. Запрос с заголовками.
// http.Get — быстрый способ, но заголовки в нём не поставить.
// Для полного контроля: http.NewRequest (собрать запрос) + client.Do (отправить).
func fetchWithHeaders(url string) error {
	// nil — тела у GET-запроса нет
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "GoCourse-Bot/1.0") // кто обращается
	req.Header.Set("Accept", "application/json")     // какой формат ответа хочу
	// req.Header.Set("Authorization", "Bearer ТОКЕН") // так передают токен доступа

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Статус:", resp.Status)
	fmt.Println("Content-Type ответа:", resp.Header.Get("Content-Type"))

	fmt.Println("Все заголовки ответа:")
	for name, values := range resp.Header {
		fmt.Printf("  %s: %v\n", name, values)
	}
	return nil
}

func main() {
	err := fetchWithHeaders("https://api.github.com/users/golang")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Хотите увидеть, что именно получил сервер?
	// Замените адрес на https://httpbin.org/headers — он вернёт наши заголовки обратно.
}
