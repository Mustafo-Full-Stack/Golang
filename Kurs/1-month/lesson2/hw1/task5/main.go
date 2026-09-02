package main

import "fmt"

func main() {

	var (
		a int
		b int
	)

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	fmt.Println("Частное:", a/b)
	fmt.Println("Остаток:", a%b)
}
