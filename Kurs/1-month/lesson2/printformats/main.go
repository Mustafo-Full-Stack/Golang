package main

import "fmt"

func main() {
	name := "Фаррух"
	age := 25
	height := 1.7854
	isStudent := true

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d лет\n", age)
	fmt.Printf("Рост: %.2f м\n", height) // 1.79 — округляет!
	fmt.Printf("Студент: %t\n", isStudent)
	fmt.Printf("Всё сразу: %s, %d, %.1f, %t\n", name, age, height, isStudent)
	fmt.Printf("Тип возраста: %T\n", age)
}
