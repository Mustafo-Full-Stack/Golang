package main

import "fmt"

func main() {
	var (
		A     int = 3
		B     int = 7
		count     = 0
	)
	for i := A; i <= B; i++ {
		fmt.Println(i)
		count++
	}
	fmt.Println("Количесво: ", count)
}
