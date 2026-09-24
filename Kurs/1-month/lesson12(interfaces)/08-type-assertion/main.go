package main

import "fmt"

func main() {
	var i any = "привет"

	// Небезопасный вариант - паника, если внутри i лежит не string.
	// s := i.(string)
	// fmt.Println(s)

	// Безопасный вариант - через ok. Программа не падает, если тип не совпал.
	// if n, ok := i.(int); ok {
	// 	fmt.Println("Это число:", n)
	// } else  {
	// 	fmt.Println("Это не число")
	// }

	//type assertion
	switch i {
	case i.(string):
		fmt.Println("это строка")
	case i.(int):
		fmt.Println("'это число'")
	case i.(bool):
		fmt.Println("'это булево параметр'")
	default:
		fmt.Println("ничего из перечисленного")
	}

	// Если раскомментировать строку ниже - будет паника, потому что
	// внутри i лежит string, а не int:
	//
	// n := i.(int)
	// panic: interface conversion: interface {} is string, not int
}
