package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// неблокирующее чтение
	select {
	case v := <-ch:
		fmt.Println("Получили:", v)
	default:
		fmt.Println("Канал пуст, идём дальше")
	}

	// неблокирующая запись
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("Записали")
	default:
		fmt.Println("Буфер полон, пропускаем")
	}

	// Практическое применение - сброс лишних событий, когда обработчик
	// не успевает их разгребать.
}
