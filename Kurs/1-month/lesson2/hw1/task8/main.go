package main

import "fmt"

const Company string = "Alif Academy"

func main() {

	var name string
	var age int
	var height float64

	var favoriteLanguage string

	var isDeveloper bool

	var projects uint

	var firstLetter rune

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Print("Введите рост: ")
	fmt.Scan(&height)

	fmt.Print("Любимый язык программирования: ")
	fmt.Scan(&favoriteLanguage)

	fmt.Print("Вы разработчик? (true/false): ")
	fmt.Scan(&isDeveloper)

	fmt.Print("Количество проектов: ")
	fmt.Scan(&projects)

	firstLetter = rune(name[0])

	fmt.Println("\n==============================")
	fmt.Println("     МОЯ ВИЗИТНАЯ КАРТОЧКА")
	fmt.Println("==============================")

	fmt.Printf("Компания: %s\n", Company)

	fmt.Printf("Имя: %s\n", name)

	fmt.Printf("Возраст: %d\n", age)

	fmt.Printf("Рост: %.2f\n", height)

	fmt.Printf("Любимый язык: %s\n", favoriteLanguage)

	fmt.Printf("Разработчик: %t\n", isDeveloper)

	fmt.Printf("Проектов создано: %d\n", projects)

	fmt.Printf("Первая буква имени: %c\n", firstLetter)

}
