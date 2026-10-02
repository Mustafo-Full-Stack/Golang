package main

import (
	"fmt"
	"time"
)

func worker(id int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Printf("Воркер %d останавливается\n", id)
			return
		default:
			fmt.Printf("Воркер %d работает\n", id)
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Отправляем сигнал остановки")
	close(done) // сигнал сразу всем
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Все остановлены")
}
