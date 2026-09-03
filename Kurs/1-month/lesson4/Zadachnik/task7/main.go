package main

import "fmt"

func main() {
	var (
		A      = 3
		B      = 6
		result = 1
	)

	for i := A; i <= B; i++ {
		result *= i
	}

	fmt.Println(result)
}
