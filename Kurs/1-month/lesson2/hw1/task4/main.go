package main

import "fmt"

func main() {

	var (
		name      string
		age       int
		height    float64
		weight    float64
		isWorking bool
	)

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Print("Введите рост: ")
	fmt.Scan(&height)

	fmt.Print("Введите вес: ")
	fmt.Scan(&weight)

	fmt.Print("Работаете? (true/false): ")
	fmt.Scan(&isWorking)

	fmt.Println("\n=== Анкета пользователя ===")

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Рост:", height)
	fmt.Println("Вес:", weight)
	fmt.Println("Работает:", isWorking)

}
