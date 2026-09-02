package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Println("Введите первый чисел : ")
	fmt.Scan(&num1)

	fmt.Println("Введите Второй чисел : ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println(num1, "Больше чем ", num2)
	}
	if num2 > num1 {
		fmt.Println(num2, "Больше чем ", num1)
	}
}
