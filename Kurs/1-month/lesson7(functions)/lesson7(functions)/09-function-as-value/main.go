package main

import "fmt"

// функцию можно передавать как обычное значение - как число или строку

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }

// op - это параметр, который сам является функцией
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	fmt.Println(apply(3, 4, add)) // 7
	fmt.Println(apply(3, 4, mul)) // 12

	// функцию можно положить в переменную
	// var operation func(int, int) int
	// operation = add
	// fmt.Println(operation(10, 5)) // 15

	// это только начало темы - подробнее про функции как значения
	// будет на одном из следующих уроков
}
