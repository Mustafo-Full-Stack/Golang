package main

import "fmt"

func main() {
	var costMoney float64 = 100

	for i := 1; i <= 10; i++ {
		kg := float64(i) / 10
		price := kg * costMoney

		fmt.Println(kg, "кг =", price)
	}
}
