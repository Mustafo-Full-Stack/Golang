package main

import (
	"fmt"
)

func main() {
	var (
		A float64
		B float64
	)

	fmt.Println("Введите любое число : ")
	fmt.Scan(&A)

	fmt.Println("ещё")
	fmt.Scan(&B)

	if A > B {
		C := A // Изменяем местах
		A = B
		B = C
	}
	fmt.Println("Новый значение А:", A)
	fmt.Println("Новый значение B:", B)
}
