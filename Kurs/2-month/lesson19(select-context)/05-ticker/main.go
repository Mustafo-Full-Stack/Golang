package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop() // обязательно, иначе утечка

	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	count := 0
	for {
		select {
		case t := <-ticker.C:
			count++
			fmt.Printf("Тик %d в %s\n", count, t.Format("15:04:05.000"))
		case <-done:
			fmt.Println("Остановлено. Всего тиков:", count)
			return
		}
	}

	// Так пишут health-check'и, сборщики метрик, периодические синхронизации.
}
