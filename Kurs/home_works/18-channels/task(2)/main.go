package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)

	go square(7, ch)
	result := <-ch

	fmt.Println("7 в квадрате =", result)
}
