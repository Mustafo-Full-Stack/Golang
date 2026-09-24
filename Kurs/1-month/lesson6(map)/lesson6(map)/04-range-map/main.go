package main

import "fmt"

func main() {
	ages := map[string]int{
		"Али":   25,
		"Вера":  30,
		"Тимур": 19,
		"Вова":  45,
	}
	//ages["Вова"] = 45

	// ключ и значение вместе
	for name, age := range ages {
		fmt.Printf("%s — %d лет\n", name, age)
	}

	// только ключи
	for name := range ages {
		fmt.Print(name, " ")
	}
	fmt.Println()

	// только значения (ключ игнорируем через _)
	for _, age := range ages {
		fmt.Print(age, " ")
	}
	fmt.Println()

	// ВАЖНО: порядок перебора случайный и меняется от запуска к запуску!
	// если нужен стабильный порядок — смотри папку sorted-output
}
