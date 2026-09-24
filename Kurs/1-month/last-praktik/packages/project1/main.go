package main

import (
	"fmt"
	"project1/User"
)

func main() {
	Hello := User.SayHello()
	fmt.Println(Hello)

	Name := User.UserName(User.User{})
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ", Name, "!")
}
