package main

import "fmt"

func main() {

	var (
		a int
		b int
	)

	a = 20
	b = 6

	fmt.Println("Первое число:", a)
	fmt.Println("Второе число:", b)

	fmt.Println("Сложение:", a+b)
	fmt.Println("Вычитание:", a-b)
	fmt.Println("Умножение:", a*b)
	fmt.Println("Деление:", a/b)
	fmt.Println("Остаток от деления:", a%b)

}
