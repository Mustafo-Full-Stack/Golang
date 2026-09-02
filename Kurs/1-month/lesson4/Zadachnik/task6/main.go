package main

import "fmt"

func main() {
	var constMoney float64 = 10

	for i := 1.2; i <= 2; i += 0.2 {
		kg := i
		price := kg * constMoney
		fmt.Printf("%.2f кг = %.2f\n", kg, price)
	}
}
