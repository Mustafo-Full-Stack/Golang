package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   string `json:"created_at"`
}

// Шаг 4. Реальный API — GitHub. Оформляем запрос в функцию,
// которая возвращает либо пользователя, либо ошибку.
func fetchGitHubUser(login string) (*GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("запрос к GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("пользователь %q не найден", login)
	}
	// 403 здесь чаще всего значит "исчерпан лимит запросов" (60 в час без токена)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub вернул статус %d", resp.StatusCode)
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("разбор ответа: %w", err)
	}
	return &user, nil
}

func main() {
	logins := []string{"golang", "torvalds", "несуществующий-юзер-12345"}

	for _, login := range logins {
		u, err := fetchGitHubUser(login)
		if err != nil {
			fmt.Println("✗", err)
			continue
		}
		fmt.Printf("✓ %s (%s): %d репозиториев, %d подписчиков\n",
			u.Name, u.Login, u.PublicRepos, u.Followers)
	}
}
