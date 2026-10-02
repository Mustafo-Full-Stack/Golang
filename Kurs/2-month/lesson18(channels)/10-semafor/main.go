package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 3) // максимум 3 одновременно
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}          // занять слот
			defer func() { <-sem }()  // освободить слот

			fmt.Printf("Задача %d выполняется\n", id)
			time.Sleep(500 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	// struct{}{} - пустая структура, занимает 0 байт.
	// Идеальна как "сигнал без данных".
}
