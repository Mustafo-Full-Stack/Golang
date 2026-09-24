package main

import "fmt"

// Урок 9, Шаг 7: замыкание-аккумулятор

func accumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	acc := accumulator()
	fmt.Println(acc(10)) // 10
	fmt.Println(acc(20)) // 30
	fmt.Println(acc(5))  // 35
}
