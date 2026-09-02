package main

import "fmt"

func main() {
	var (
		cnt      = 0
		number_1 int
		number_2 int
		number_3 int
	)

	fmt.Println("Введите Первый число : ")
	fmt.Scan(&number_1)

	fmt.Println("Введите Второй число : ")
	fmt.Scan(&number_2)

	fmt.Println("Введите Третий число : ")
	fmt.Scan(&number_3)

	if number_1 > 0 {
		cnt++
	}
	if number_2 > 0 {
		cnt++
	}
	if number_3 > 0 {
		cnt++
	}

	fmt.Println("Положительные числи :", cnt)
}
