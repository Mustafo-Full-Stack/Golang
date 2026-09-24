package main

import "fmt"

// Урок 9, Шаг 2: немедленный вызов анонимной функции - () сразу после тела

func main() {
	// func() {
	// 	fmt.Println("Меня вызвали сразу")
	// }()

	// result := func(a, b int) int {
	// 	return a * b
	// }(3, 4)
	// fmt.Println(result) // 12

	//fmt.Println(greet("Bob"))

	customGreet := func(name1, name2 string) string {
		return "Hello " + name1 + name2
	}

	fmt.Println(customGreet("Bob","Alex"))

	//fmt.Printf("%T\n", customGreet)
}

func greet(name string) string {
	return "Hello " + name
}
