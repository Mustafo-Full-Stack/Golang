package main

import (
	"fmt"
	"sync"
)

// RWMutex - оптимизация для случая "читают намного чаще, чем пишут".
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // много читателей одновременно - можно
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // писатель - только один, и без читателей
	defer c.mu.Unlock()
	c.data[key] = value
}

func main() {
	cache := &Cache{data: make(map[string]string)}
	cache.Set("name", "Го")

	v, ok := cache.Get("name")
	fmt.Println(v, ok)

	// Если данные читают в 100 раз чаще, чем пишут, RWMutex даёт
	// заметный выигрыш по сравнению с обычным Mutex.
}
