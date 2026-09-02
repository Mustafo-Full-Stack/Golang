package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите Число: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println(number + 1)
	} else {
		fmt.Println(number - 2)
	}
}
