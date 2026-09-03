package main

import "fmt"

func main() {
	var (
		SecretNumber int = 37
		UserNumber   int
	)

	fmt.Println("Введите Секретное Число: ")
	fmt.Scan(&UserNumber)

	if UserNumber > SecretNumber {
		fmt.Println("Ваша Число больше чем Секретное число !")
	} else if UserNumber < SecretNumber {
		fmt.Println("Ваша Число меньше чем Секретное число !")
	} else if UserNumber == SecretNumber {
		return
	}
}
