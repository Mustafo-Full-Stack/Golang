package main

import "fmt"

// факториал считаем циклом: n! = 1 * 2 * 3 * ... * n
func factorial(n int) (result int) {
	result = 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return
}

func main() {
	var n int
	fmt.Print("Введите число n:")
	fmt.Scanln(&n)

	result := factorial(n)
	fmt.Println(result)
}
