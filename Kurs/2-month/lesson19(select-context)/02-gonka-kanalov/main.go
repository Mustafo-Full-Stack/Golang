package main

import (
	"fmt"
	"time"
)

func fetchFrom(source string, delay time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(delay)
		ch <- fmt.Sprintf("данные от %s", source)
	}()
	return ch
}

func main() {
	server1 := fetchFrom("сервер-1", 300*time.Millisecond)
	server2 := fetchFrom("сервер-2", 150*time.Millisecond)
	server3 := fetchFrom("сервер-3", 500*time.Millisecond)

	select {
	case r := <-server1:
		fmt.Println("Первым ответил:", r)
	case r := <-server2:
		fmt.Println("Первым ответил:", r)
	case r := <-server3:
		fmt.Println("Первым ответил:", r)
	}

	// Реальный приём: запросили у трёх реплик, взяли самый быстрый ответ.
}
