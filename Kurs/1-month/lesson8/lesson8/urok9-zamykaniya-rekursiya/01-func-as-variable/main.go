package main

import "fmt"

// Урок 9, Шаг 1: функцию можно положить в переменную

func main() {
	square := func(x int) int {
		return x * x
	}

	fmt.Println(square(5))     // 25
	fmt.Printf("%T\n", square) // func(int) int
}
