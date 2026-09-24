package main

import (
	"fmt"
	"strings"
)

// Урок 9, Шаг 9: рекурсия - факториал (+трассировка вызовов)

func factorial(n int) int {
	if n <= 1 {
		return 1 // базовый случай
	}
	return n * factorial(n-1) // рекурсивный случай
}

func factorialTrace(n int, depth int) int {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%sВход factorial(%d)\n", indent, n)
	if n <= 1 {
		fmt.Printf("%sБазовый случай, возврат 1\n", indent)
		return 1
	}
	res := n * factorialTrace(n-1, depth+1)
	fmt.Printf("%sВозврат %d\n", indent, res)
	return res
}

func main() {
	fmt.Println(factorial(5)) // 120

	// раскрутка на доске:
	// factorial(5) = 5 * factorial(4)
	//              = 5 * 4 * factorial(3)
	//              = 5 * 4 * 3 * factorial(2)
	//              = 5 * 4 * 3 * 2 * factorial(1)
	//              = 5 * 4 * 3 * 2 * 1 = 120

	factorialTrace(5, 0) // визуализация вызовов помогает понять рекурсию
}
