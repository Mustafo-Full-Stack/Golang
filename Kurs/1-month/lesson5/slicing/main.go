package main

import "fmt"

func main() {
	data := []int{10, 20, 30, 40, 50}
	fmt.Println("До:", data)

	// slicing
	// правильно подсчета диапазона элементов
	// [i:j] - от i индекс до j-1

	dataAnother := data[1:4]
	fmt.Println(dataAnother) // [20 30 40] — с 1 по 3
	// len = 3 cap = 4
	data[0] = 5500
	data[1] = 199
	data[3] = 222
	fmt.Println("После:", data)
	// fmt.Println(data[:3])  // [10 20 30] — с начала
	// fmt.Println(data[2:])  // [30 40 50] — до конца
	// fmt.Println(data[:])   // весь срез
	// fmt.Println(data[2:3]) // [30] — один элемент, но это срез
}
