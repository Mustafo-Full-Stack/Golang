package main

import (
	"fmt"

	"task2/client"
)

func main() {
	user, err := client.GetClient(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(user.Name, user.Email, user.Phone)
}
