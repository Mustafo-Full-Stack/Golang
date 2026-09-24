package main

import "fmt"

func main() {
	var CityS string = "Москва"

	fmt.Println("ДО:", CityS)

	pointer := &CityS
	*pointer = "Душанбе"

	fmt.Println("ПОСЛЕ:", CityS)
}
