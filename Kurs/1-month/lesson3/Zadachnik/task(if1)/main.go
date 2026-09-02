package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число :")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println(num + 1)
	} else {
		fmt.Println(num)
	}
}
