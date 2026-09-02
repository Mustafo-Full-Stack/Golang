package main

import "fmt"

func main() {

	var name string = "Ivan"
	var surname string = "Ivanov"

	var age int = 31
	var experience int = 3

	var salary float64 = 5000.50

	var favoriteLanguage string = "Go"

	var commercialExperience bool = true

	var knowsDocker bool = true

	var firstLetter rune = 'R'

	var favoriteSymbol byte = '@'

	var projects uint = 10

	fmt.Printf("=== Паспорт разработчика ===\n\n")

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Фамилия: %s\n", surname)

	fmt.Printf("Возраст: %d\n", age)

	fmt.Printf("Опыт: %d года\n", experience)

	fmt.Printf("Зарплата: %.2f$\n", salary)

	fmt.Printf("Любимый язык: %s\n", favoriteLanguage)

	fmt.Printf("Коммерческий опыт: %t\n", commercialExperience)

	fmt.Printf("Знает Docker: %t\n", knowsDocker)

	fmt.Printf("Первая буква имени: %c\n", firstLetter)

	fmt.Printf("Любимый символ: %c\n", favoriteSymbol)

	fmt.Printf("Количество проектов: %d\n", projects)

}
