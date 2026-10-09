package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(string(body))

	fmt.Println("Заголовок:", resp.Header.Get("Content-Type"))
	fmt.Println("Статус код:", resp.StatusCode)
}
