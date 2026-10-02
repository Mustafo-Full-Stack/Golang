package main

import "fmt"

// producer - функция которая отправляет данные в канал
func producer(ch chan<- int, n int) { 
	for i := 1; i <= n; i++ {
		fmt.Println("Отправляю", i)
		ch <- i
	}
	close(ch) // ОБЯЗАТЕЛЬНО, иначе range зависнет навсегда
}

func main() {
	ch := make(chan int)
	go producer(ch, 5)

	for v := range ch {
		fmt.Println("Получил", v)
	}
	fmt.Println("Канал закрыт, цикл завершён")

	// Убери close(ch) в producer и покажи deadlock - очень поучительно.
}
