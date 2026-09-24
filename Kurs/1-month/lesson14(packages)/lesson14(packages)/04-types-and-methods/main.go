package main

import (
	"fmt"

	"myapp/models"
)

func main() {
	u, err := models.New(1, "Али", "ali@mail.tj")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := u.SetPassword("supersecret"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)
	fmt.Println(u.CheckPassword("wrong"))       // false
	fmt.Println(u.CheckPassword("supersecret")) // true

	_, err = models.New(2, "", "bad-email")
	fmt.Println("ошибка при пустом имени:", err)

	// Поломка: раскомментируй - u.password не экспортировано,
	// снаружи пакета models к нему нет доступа.
	// u.password = "hack"
}
