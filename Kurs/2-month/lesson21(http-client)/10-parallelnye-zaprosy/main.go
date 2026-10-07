package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
}

func fetchGitHubUser(login string) (*GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("запрос к GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("пользователь %q не найден", login)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub вернул статус %d", resp.StatusCode)
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("разбор ответа: %w", err)
	}
	return &user, nil
}

// Результат одного запроса: либо данные, либо ошибка
type Result struct {
	Login     string
	Followers int
	Err       error
	Duration  time.Duration
}

// Шаг 10. Параллельные запросы: горутины + WaitGroup из уроков 17–18.
func main() {
	logins := []string{"golang", "torvalds", "rsc", "bradfitz", "robpike"}

	// Способ 1: последовательно — каждый запрос ждёт предыдущий
	start := time.Now()
	for _, login := range logins {
		fetchGitHubUser(login)
	}
	seqTime := time.Since(start)
	fmt.Println("Последовательно:", seqTime.Round(time.Millisecond))

	// Способ 2: параллельно — все запросы летят одновременно
	start = time.Now()
	results := make([]Result, len(logins))
	var wg sync.WaitGroup

	for i, login := range logins {
		wg.Add(1)
		go func(idx int, login string) {
			defer wg.Done()

			t := time.Now()
			u, err := fetchGitHubUser(login)

			r := Result{Login: login, Err: err, Duration: time.Since(t)}
			if err == nil {
				r.Followers = u.Followers
			}
			// Пишем в СВОЮ ячейку results[idx] — гонки нет, мьютекс не нужен
			results[idx] = r
		}(i, login)
	}
	wg.Wait()
	parTime := time.Since(start)

	fmt.Println("Параллельно:    ", parTime.Round(time.Millisecond))
	fmt.Printf("Ускорение: %.1fx\n\n", float64(seqTime)/float64(parTime))

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%-10s ошибка: %v\n", r.Login, r.Err)
			continue
		}
		fmt.Printf("%-10s %6d подписчиков (%v)\n",
			r.Login, r.Followers, r.Duration.Round(time.Millisecond))
	}
}
