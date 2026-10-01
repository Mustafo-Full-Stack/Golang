package main

import (
	"fmt"
	"sync"
)

// Решение гонки данных - sync.Mutex ("замок" на данные).
func simpleCounter() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // пропускает внутрь только одну горутину
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Результат:", counter) // всегда 1000
}

// Более практично - мьютекс внутри структуры.
type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{counts: make(map[string]int)}
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}

func main() {
	simpleCounter()

	c := NewSafeCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("visits")
		}()
	}
	wg.Wait()
	fmt.Println("Посещений:", c.Value("visits")) // 100

	// Важно: обычная map НЕ потокобезопасна. Запись из нескольких горутин
	// без мьютекса может вызвать панику "concurrent map writes".
}
