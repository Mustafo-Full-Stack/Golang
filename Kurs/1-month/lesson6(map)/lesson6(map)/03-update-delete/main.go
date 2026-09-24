package main

import "fmt"

func main() {
	ages := map[string]int{
		"Али":   25,
		"Вера":  30,
		"Тимур": 19,
	}

	ages["Али"] = 26
	ages["Новый"] = 40

	fmt.Printf("наша map:%v\n", ages)

	fmt.Println("До проверки ключа Несуществующий")

	_, ok := ages["Вера"]
	if ok {
		delete(ages, "Вера") // ничего не происходит
		fmt.Println(ages, len(ages))
		fmt.Println("После проверки ключа Вера")
	} else {
		fmt.Println("увы нет ключа!")
	}
}
