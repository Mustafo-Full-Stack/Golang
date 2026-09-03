package main

import "fmt"

func main() {
	var (
		N     int
		Error string = "Не найдено"
	)
	fmt.Println("Введите число N:")
	fmt.Scan(&N)

	for i := 1; i < N; i++ {
		if N%3 == 0 && N%7 == 0 {
			fmt.Println(N)
			break
		} else {
			fmt.Println(Error)
			break
		}
	}
}
