package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

// Шаг 7. Query-параметры: то, что идёт в адресе после знака "?".
// Например: .../repositories?q=golang&per_page=5
func search(query string, limit int) error {
	base, err := url.Parse("https://api.github.com/search/repositories")
	if err != nil {
		return err
	}

	// url.Values собирает параметры и сам экранирует спецсимволы
	params := url.Values{}
	params.Add("q", query)
	params.Add("per_page", fmt.Sprint(limit))
	params.Add("sort", "stars")
	base.RawQuery = params.Encode()

	fmt.Println("URL:", base.String())

	resp, err := client.Get(base.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub вернул статус %d", resp.StatusCode)
	}

	// Структура нужна только здесь, один раз — поэтому анонимная (урок про структуры)
	var result struct {
		TotalCount int `json:"total_count"`
		Items      []struct {
			FullName string `json:"full_name"`
			Stars    int    `json:"stargazers_count"`
		} `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Printf("Найдено репозиториев: %d\n", result.TotalCount)
	for i, r := range result.Items {
		fmt.Printf("%d. %-40s ★%d\n", i+1, r.FullName, r.Stars)
	}
	return nil
}

func main() {
	// Зачем url.Values? Смотрим, как он экранирует пробелы, кириллицу и "&"
	demo := url.Values{}
	demo.Add("q", "привет мир & go")
	fmt.Println("Экранирование:", demo.Encode())
	fmt.Println()

	if err := search("golang", 5); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
