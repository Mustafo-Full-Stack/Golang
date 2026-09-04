package main

import "fmt"

func main() {
	data := []int{10, 20, 30, 40, 50}

	for i, v := range data {
		fmt.Printf("Индекс %d: значение %d\n", i, v)
	}

	// только значения
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// только индексы
	for i := range data {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//  Срезы делят память
	original := []int{1, 2, 3, 4, 5}
	fmt.Printf("original: len = %d, cap = %d", len(original), cap(original))
	part := original[1:3]
	fmt.Printf("part: len = %d, cap = %d", len(part), cap(part))

	part[0] = 999
	fmt.Println(part)
	fmt.Println(original)
}
