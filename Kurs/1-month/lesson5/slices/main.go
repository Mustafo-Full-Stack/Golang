package main

import "fmt"

/* slices(срез) - это динамический массив, который может изменять свой размер
во время выполнения программы.
Срезы предоставляют более гибкий способ работы с
последовательностями данных по сравнению с массивами.*/

func main() {
	fruits := []string{"яблоко", "банан"}         // len = 2 cap = 2
	fmt.Println(fruits, len(fruits), cap(fruits)) // [яблоко банан] len = 2 cap = 2

	fruits = append(fruits, "вишня", "персик", "черника")

	fmt.Println(fruits, len(fruits), cap(fruits))
	// [яблоко банан вишня] 3 4

	fruits = append(fruits, "дыня", "груша")
	fmt.Println(fruits, len(fruits), cap(fruits))

	// рост среза
	// var s []int
	// prevCap := cap(s)
	// fmt.Printf("len=%2d  cap=%d\n", len(s), cap(s))
	// for i := 1; i <= 20; i++ {
	// 	s = append(s, i)
	// 	if cap(s) != prevCap {
	// 		fmt.Printf("len=%2d  cap вырос до %d\n", len(s), cap(s))
	// 		prevCap = cap(s)
	// 	}
	// }
}
