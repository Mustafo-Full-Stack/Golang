package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		// time.Sleep(100 * time.Millisecond)
		ch4 <- "сообщение из ch4"
	}()

	go func() {
		// time.Sleep(200 * time.Millisecond)
		ch1 <- "сообщение из ch1"
	}()

	go func() {
		// time.Sleep(100 * time.Millisecond)
		ch3 <- "сообщение из ch3"
	}()

	go func() {
		// time.Sleep(100 * time.Millisecond)
		ch2 <- "сообщение из ch2"
	}()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg := <-ch1:
	// 		fmt.Println("Получено:", msg)
	// 	case msg := <-ch2:
	// 		fmt.Println("Получено:", msg)
	// 	}
	// }

	select {
	case msg := <-ch1:
		fmt.Println("Получено:", msg)
	case msg := <-ch4:
		fmt.Println("Получено:", msg)
	case msg := <-ch2:
		fmt.Println("Получено:", msg)
	case msg := <-ch3:
		fmt.Println("Получено:", msg)

	}
}
