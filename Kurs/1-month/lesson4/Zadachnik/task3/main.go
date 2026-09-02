package main

import "fmt"

func main() {
	var (
		A     int = 3
		B     int = 8
		count     = 0
	)
	for i := B; i > A; i-- { // i = 8; 8 > 3; i--
		fmt.Println(i)
		count++
	}
	fmt.Println("Количесво: ", count)
}
