package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(name string, wg *sync.WaitGroup) {
	defer wg.Done() // сообщаем: "я закончил"

	fmt.Printf("Начали %s\n", name)
	time.Sleep(1 * time.Second) // иммитация сети
	fmt.Printf("Готово %s\n", name)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	files := []string{"file1.zip", "file2.zip", "file3.zip", "file4.zip", "file5.zip"}

	for _, f := range files {
		wg.Add(1) // ВАЖНО: Add вызываем ДО запуска горутины
		go downloadFile(f, &wg)
	}

	wg.Wait() // блокируемся, пока счётчик не станет 0
	fmt.Println("Всего заняло:", time.Since(start)) // ~1 секунда для 5 файлов!

	// Три правила WaitGroup:
	// 1. Add(1) - до go
	// 2. Done() - через defer, первой строкой
	// 3. WaitGroup передаём указателем (*sync.WaitGroup)
}
