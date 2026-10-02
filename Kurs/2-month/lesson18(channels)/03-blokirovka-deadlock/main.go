package main

import "fmt"

func main() {
	ch := make(chan int)

	// Сначала покажи поломанный вариант (раскомментируй, чтобы увидеть deadlock):
	// ch <- 42 // блокируется навсегда - никто не читает
	// fmt.Println(<-ch)
	//
	// fatal error: all goroutines are asleep - deadlock!
	// Go честно говорит, что программа зависла бы навсегда, и сам её останавливает.

	// Исправление - отправлять из отдельной горутины:
	go func() { ch <- 42 }()
	fmt.Println(<-ch)
}
