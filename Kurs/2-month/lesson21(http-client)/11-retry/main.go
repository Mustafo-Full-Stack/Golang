package main

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

// Шаг 11. Retry: если запрос не удался, пробуем ещё раз.
// С каждой попыткой ждём вдвое дольше (200мс, 400мс, 800мс...) —
// так мы не долбим упавший сервер и даём ему время прийти в себя.
func fetchWithRetry(url string, attempts int) (*http.Response, error) {
	var lastErr error
	delay := 200 * time.Millisecond

	for i := 1; i <= attempts; i++ {
		resp, err := client.Get(url)

		// Успех: запрос прошёл и это не 5xx.
		// 4xx не повторяем — там виноваты мы, второй раз ответ не изменится.
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}

		if err != nil {
			lastErr = err
		} else {
			resp.Body.Close() // эту попытку выбрасываем — тело закрываем
			lastErr = fmt.Errorf("статус %d", resp.StatusCode)
		}

		fmt.Printf("Попытка %d из %d не удалась (%v)\n", i, attempts, lastErr)

		// После последней попытки ждать незачем
		if i < attempts {
			fmt.Println("  жду", delay)
			time.Sleep(delay)
			delay *= 2
		}
	}
	return nil, fmt.Errorf("после %d попыток: %w", attempts, lastErr)
}

func main() {
	// Хороший адрес: ответ с первой попытки
	resp, err := fetchWithRetry("https://jsonplaceholder.typicode.com/todos/1", 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Успех:", resp.Status)
		resp.Body.Close()
	}
	fmt.Println()

	// Плохой адрес: httpbin всегда отвечает 503, все попытки провалятся
	resp, err = fetchWithRetry("https://httpbin.org/status/503", 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Успех:", resp.Status)
		resp.Body.Close()
	}
}
