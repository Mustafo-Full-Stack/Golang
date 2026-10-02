package main

import (
	"fmt"
	"time"
)

func worker(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Получен сигнал остановки")
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})
	go worker(done)

	time.Sleep(1 * time.Second)
	close(done) // закрытие канала = сигнал СРАЗУ ВСЕМ получателям
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main завершён")

	// Здесь впервые появляется select - завтра разберём его подробно.
}
