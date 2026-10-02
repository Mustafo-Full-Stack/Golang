package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)

	v1, ok1 := <-ch
	fmt.Println(v1, ok1) // 1 true
	v2, ok2 := <-ch
	fmt.Println(v2, ok2) // 2 true
	v3, ok3 := <-ch
	fmt.Println(v3, ok3) // 0 false - канал закрыт и пуст

	// Та же идиома comma-ok, что и при чтении из map.
	// Правило: закрывает всегда отправитель, никогда получатель.
}
