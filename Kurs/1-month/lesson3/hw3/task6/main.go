package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число : ")
	fmt.Scan(&a)
	count := 0

	for i := a; i > 0; i /= 10 {
		if i%10 == 0 {
			count++
		}
	}
	fmt.Println("Количество нулей:", count)
}
