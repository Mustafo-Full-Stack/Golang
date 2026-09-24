package main

import "fmt"

// число простое, если оно больше 1 и делится только на 1 и на само себя
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число n:")
	fmt.Scanln(&n)

	result := isPrime(n)
	fmt.Println(result)
}
