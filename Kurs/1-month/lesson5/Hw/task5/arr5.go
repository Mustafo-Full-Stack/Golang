package main

import "fmt"

func main() {
	var zarb int = 1
	nums := []int{1, 2, 3, 4, 5}
	for _, nums := range nums {
		zarb *= nums
	}
	fmt.Println("Результать Умножениеe на 2 :", zarb)
}
