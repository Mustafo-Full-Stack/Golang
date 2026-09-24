package main

import "fmt"

// функция может ничего не возвращать - тогда после списка параметров
// просто нет типа результата и в теле нет return со значением

// Ali
func greet(name string) {
	fmt.Println("Привет,", name+"!")
}

func customGreet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}

func printLine() {
	fmt.Println("------------------------")
}

func main() {
	fmt.Println(customGreet("Вера"))
	//customGreet("Вера")
}
