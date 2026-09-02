package main

import "fmt"

func main() {
	number := 0

	fmt.Println("Введите Число: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println(number + 1)
	} else if number < 0 {
		fmt.Println(number - 2)
	} else if number == 0 {
		fmt.Println(number + 10)
	}
}
