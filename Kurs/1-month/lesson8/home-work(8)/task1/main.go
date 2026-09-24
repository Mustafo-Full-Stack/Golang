package main

import "fmt"

func main() {
	age := 10
	fmt.Println("Значение:", age)
	fmt.Println("Адрес:", &age)
	fmt.Printf("Тип адрес: %T", &age)
}
