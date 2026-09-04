package main

import "fmt"

func main() {
	var sum int = 0

	nums := []int{1, 2, 3, 4}
	nums = append(nums, 5)
	for _, numResult := range nums {
		sum = sum + numResult
	}
	avarage := sum / len(nums)
	fmt.Println("Summa :", sum)
	fmt.Println("Среднее Число :", avarage)
}
