/*
For11. Дано целое число N (N > 0). Найти сумму:
N² + (N + 1)² + (N + 2)² + ... + (2·N)².
Результат — целое число.
*/

package main

import "fmt"

func main() {
	var (
		N      = 5
		result = 0
	)
	for i := N; i <= 2*N; i++ { // 5 <= 10
		result += i * i // 5*5 + 6*6 + 7*7 + 8*8 + 9*9 + 10*10 = 355
	}
	fmt.Println(result)
}
