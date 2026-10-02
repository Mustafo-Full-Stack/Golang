package main

import (
	"fmt"
	"time"
)

func slowOperation(d time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(d)
		ch <- "операция завершена"
	}()
	return ch
}

func main() {
	// успеваем
	select {
	case res := <-slowOperation(500 * time.Millisecond):
		fmt.Println("OK:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Таймаут!")
	}

	// не успеваем
	select {
	case res := <-slowOperation(2 * time.Second):
		fmt.Println("OK:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Таймаут! Операция слишком долгая")
	}

	// В любом продакшн-сервисе каждый внешний запрос обёрнут таймаутом.
	// Без этого сервис умирает, когда тормозит партнёрский API.
}
