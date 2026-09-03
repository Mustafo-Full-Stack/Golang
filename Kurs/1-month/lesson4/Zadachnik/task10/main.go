//For10. Дано целое число N (N > 0). Найти сумму:
//1 + 1/2 + 1/3 + ... + 1/N.
//Результат вывести как вещественное число.

package main

import "fmt"

func main() {
	var (
		N      = 5
		result = 0.0
	)
	for i := 1; i <= N; i++ {
		result += 1 / float64(i)
	}
	fmt.Println(result)
}
