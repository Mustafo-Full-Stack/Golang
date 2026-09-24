package main

import "fmt"

// Урок 10, Шаг 1: проблема - данные одного человека разложены по разным переменным

func main() {
	name1, age1, city1 := "Али", 25, "Душанбе"
	name2, age2, city2 := "Вера", 30, "Худжанд"

	fmt.Println(name1, age1, city1) // Али 25 Душанбе
	fmt.Println(name2, age2, city2) // Вера 30 Худжанд
}

