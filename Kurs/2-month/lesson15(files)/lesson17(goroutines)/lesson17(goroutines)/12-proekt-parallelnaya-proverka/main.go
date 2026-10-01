package main

import (
	"fmt"
	"sync"
	"time"
)

// Итоговый проект урока: параллельная проверка списка сайтов.
// Реалистичный сценарий - на следующих уроках заменим имитацию
// настоящими HTTP-запросами.

type Result struct {
	URL    string
	Status int
	Time   time.Duration
}

func checkSite(url string, wg *sync.WaitGroup, mu *sync.Mutex, results *[]Result) {
	defer wg.Done()

	start := time.Now()
	time.Sleep(time.Duration(len(url)*20) * time.Millisecond) // имитация запроса
	elapsed := time.Since(start)

	mu.Lock()
	*results = append(*results, Result{URL: url, Status: 200, Time: elapsed})
	mu.Unlock()
}

func main() {
	urls := []string{
		"google.com", "github.com", "stackoverflow.com",
		"golang.org", "wikipedia.org", "reddit.com",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]Result, 0, len(urls))

	start := time.Now()
	for _, u := range urls {
		wg.Add(1)
		go checkSite(u, &wg, &mu, &results)
	}
	wg.Wait()

	fmt.Printf("Проверено %d сайтов за %v\n\n", len(results), time.Since(start))
	for _, r := range results {
		fmt.Printf("%-20s %d  %v\n", r.URL, r.Status, r.Time.Round(time.Millisecond))
	}
}
