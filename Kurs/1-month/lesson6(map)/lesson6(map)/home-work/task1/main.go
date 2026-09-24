package main

import "fmt"

func main() {
	var name string

	var PhoneNum = map[string]int{
		"Ali":  900000001,
		"Bob":  900000002,
		"John": 900000003,
		"Mike": 900000004,
		"Sara": 900000005,
	}

	fmt.Println("Введите Имя:")
	fmt.Scanln(&name)
	number, ok := PhoneNum[name]

	if ok {
		fmt.Println("Контакт найден")
		fmt.Println("Номер:", number)
	} else {
		fmt.Println("Контакт Ненайден")
	}
	// Ali   → 900000001
	// Bob   → 900000002
	// John  → 900000003
	// Mike  → 900000004
	// Sara  → 900000005
}
