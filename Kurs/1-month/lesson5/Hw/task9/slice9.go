package main

import "fmt"

func main() {
	var Positive = 0
	var Negative = 0
	var Number = []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(Number)-1; i++ {
		if Number[i] > 0 {
			if Number[i]%2 == 0 {
				Positive++
			} else {
				Negative++
			}
		}
	}
	fmt.Println("Чётный чисел :", Positive)
	fmt.Println("Нечётный чисел:", Negative)
}
