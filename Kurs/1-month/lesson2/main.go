package main

import "fmt"

func main() {

	// ===== Целые числа =====
	// var название-переменной тип-переменной = значение
	var age int = 31
	// camelCase (myDog, myCar)
	// no go standart: snakeCase(my_car, short_form_age)
	//kebabCase (short-form-age)
	//shortFormAge := 31
	var temperature int8 = -15
	//shortFormTemp := -15
	var population int64 = 8_200_000_000

	// ===== Беззнаковые числа =====
	//var gameLevel uint = 10
	// gameLever := 10
	var score uint32 = 5000

	// ===== Числа с плавающей точкой =====
	var price float32 = 99.99
	var pi float64 = 3.1415926535

	// ===== Комплексные числа =====
	var c1 complex64 = 2 + 3i
	var c2 complex128 = 10 + 5i

	// ===== Логический тип =====
	var isAdmin bool = true
	// isAdmin := false

	// ===== Строка =====
	var name string = "Bob"

	// ===== Руны (символ Unicode) =====
	var letter rune = 'Ж'

	// ===== Байт =====
	var symbol byte = 'A'

	fmt.Printf("age         = %v, type = %T\n", age, age)
	fmt.Printf("temperature = %v, type = %T\n", temperature, temperature)
	fmt.Printf("population  = %v, type = %T\n", population, population)

	fmt.Println()

	fmt.Printf("level = %v, type = %T\n", level, level)
	fmt.Printf("score = %v, type = %T\n", score, score)

	fmt.Println() // new line

	fmt.Printf("price = %v, type = %T\n", price, price)
	fmt.Printf("pi    = %v, type = %T\n", pi, pi)

	fmt.Println()

	fmt.Printf("c1 = %v, type = %T\n", c1, c1)
	fmt.Printf("c2 = %v, type = %T\n", c2, c2)

	fmt.Println()

	fmt.Printf("isAdmin = %v, type = %T\n", isAdmin, isAdmin)
	fmt.Printf("name    = %v, type = %T\n", name, name)
	fmt.Printf("letter  = %v, type = %T\n", letter, letter)
	fmt.Printf("symbol  = %v, type = %T\n", symbol, symbol)
}
