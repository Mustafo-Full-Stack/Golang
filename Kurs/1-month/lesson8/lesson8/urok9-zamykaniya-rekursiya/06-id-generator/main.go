package main

import "fmt"

// Урок 9, Шаг 6: практичное замыкание - генератор ID

func idGenerator(prefix string) func() string {
	id := 0
	return func() string {
		id++
		return fmt.Sprintf("%s-%04d", prefix, id)
	}
}

func main() {
	userID := idGenerator("USER")
	orderID := idGenerator("ORDER")

	fmt.Println(userID())  // USER-0001
	fmt.Println(userID())  // USER-0002
	fmt.Println(orderID()) // ORDER-0001
}
