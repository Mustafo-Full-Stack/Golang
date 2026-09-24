package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Print("Введите число n:")
	fmt.Scanln(&n)

	result := isEven(n)
	fmt.Println(result)
}
