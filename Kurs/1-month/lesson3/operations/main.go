package main

import "fmt"

func main() {
	a, b := 7, 3

	fmt.Println("Сумма:      ", a+b)
	fmt.Println("Разность:   ", a-b)
	fmt.Println("Произведение:", a*b)
	fmt.Println("Деление:    ", a/b) // целочисленное 
	fmt.Println("Остаток:    ", a%b) // деление с остатком

	seconds := 3725
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60
	fmt.Printf("%v сек = %v ч %v мин %v сек\n", seconds, hours, minutes, secs)
}
