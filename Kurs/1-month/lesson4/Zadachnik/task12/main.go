package main

import "fmt"

func main() {
	var N int = 5
	result := 1.0

	for i := 1; i <= N; i++ {
		m := 1.0 + float64(i)/10 // 1.0 + i/10, и другие крч..  до 3.6036
		result *= m
	}
	fmt.Println(result)
}
