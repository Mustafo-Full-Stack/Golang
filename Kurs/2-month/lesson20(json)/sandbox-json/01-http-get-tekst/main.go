package main

import (
	"fmt"
	"io"
	"net/http"
)

// Шаг 1. Просто ходим в интернет и смотрим, что вернулось.
// Это "сырой" JSON-текст, пока мы его не разбираем.
func main() {
	// Бесплатный тестовый API: отдаёт данные про вымышленного пользователя
	url := "https://jsonplaceholder.typicode.com/users/1"

	// http.Get — отправить запрос GET (попросить данные)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	// Тело ответа нужно ЗАКРЫТЬ, как файл (урок про файлы)
	defer resp.Body.Close()

	fmt.Println("Статус:", resp.Status) // 200 OK

	// Body — это io.Reader, читаем всё в []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Println(string(body))
}
