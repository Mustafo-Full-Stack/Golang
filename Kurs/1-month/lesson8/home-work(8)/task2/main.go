package main

import "fmt"

func main() {
	var temperatue float64 = 25.0

	fmt.Println("ДО :", temperatue)
	ptr := &temperatue
	*ptr = 5.0
	fmt.Println("ПОСЛЕ :", temperatue)
}
