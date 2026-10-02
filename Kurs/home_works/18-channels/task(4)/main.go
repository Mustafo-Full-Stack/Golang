package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("ПОСЛЕ ЗАПОЛЬНЕНИЕ:")
	fmt.Println(len(ch), cap(ch))

	v1 := <-ch
	v2 := <-ch
	v3 := <-ch

	fmt.Println("ПОСЛЕ ЧТЕНИЕ:")
	fmt.Println(len(ch), cap(ch))

	fmt.Println("Значение всех:", v1, v2, v3)
}
