package main

import "fmt"

func main() {
	var (
		CandyKG   int = 10
		CostMoney int = 100
	)

	for i := 1; i <= CandyKG; i++ {
		fmt.Println(i * int(CostMoney))
	}
}
