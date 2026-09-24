package main

import "fmt"

// "..." значит "любое количество аргументов" - внутри функции
// nums превращается в обычный срез float64
func average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0 // чисел нет - делить не на что
	}

	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func main() {
	fmt.Println(average(2, 4, 6))       // 4
	fmt.Println(average(1, 2, 3, 4, 5)) // 3
	fmt.Println(average())              // 0
}
