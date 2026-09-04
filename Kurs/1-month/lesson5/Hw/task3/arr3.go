package main

import "fmt"

func main() {
	var sum int = 0
	nums := []int{1, 2, 3, 4, 5, 6}
	for _, nums := range nums {
		sum = sum + nums
	}
	fmt.Println("Сумма :", sum)
}
