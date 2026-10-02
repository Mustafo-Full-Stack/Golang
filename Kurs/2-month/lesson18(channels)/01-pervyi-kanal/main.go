package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Горутина: отправляю сообщение")
		ch <- "Привет из горутины!"
		fmt.Println("Горутина: сообщение отправлено")
	}()

	fmt.Println("Main: жду сообщение...")
	msg := <-ch
	fmt.Println("Main: получил:", msg)

	// Обрати внимание на порядок вывода - он показывает синхронизацию.
	// time.Sleep больше не нужен - канал сам заставил main подождать!
}
