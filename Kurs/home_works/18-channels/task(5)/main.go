package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40

	fmt.Println("ПОСЛЕ ЗАПОЛЬНЕНИЕ:")
	fmt.Println(len(ch), cap(ch))

	v1 := <-ch
	v2 := <-ch
	v3 := <-ch

	fmt.Println("ПОСЛЕ ЧТЕНИЕ:")
	fmt.Println(len(ch), cap(ch))

	fmt.Println("Значение всех:", v1, v2, v3)
	//deadlock, потому что у нас
	//там зафиксированный размер как массив
	//и место нету хранить 40 чтобы
	//продолжить
}
