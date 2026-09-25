package main

import (
	"fmt"
	"project3/users"
)

// TODO: самостоятельно написать сценарий программы.

func main() {
	user1, err := users.GetUser("", "test@gmail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(user1)
}
