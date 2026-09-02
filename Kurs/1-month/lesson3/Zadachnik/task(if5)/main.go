package main

import "fmt"

func main() {
	var (
		positiveCnt = 0 //Сеенти(Cnt) Это счётик
		negativeCnt = 0
		number_1    int
		number_2    int
		number_3    int
	)

	fmt.Println("Введите первое число: ")
	fmt.Scan(&number_1)

	fmt.Println("Введите второе число: ")
	fmt.Scan(&number_2)

	fmt.Println("Введите третие число: ")
	fmt.Scan(&number_3)

	if number_1 > 0 {
		positiveCnt++
	} else if number_1 < 0 {
		negativeCnt++
	}

	if number_2 > 0 {
		positiveCnt++
	} else if number_2 < 0 {
		negativeCnt++
	}

	if number_3 > 0 {
		positiveCnt++
	} else if number_3 < 0 {
		negativeCnt++
	}
	fmt.Println("Положительных чисел:", positiveCnt)
	fmt.Println("Отрицательных чисел:", negativeCnt)
}
