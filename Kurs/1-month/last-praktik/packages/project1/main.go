package main

import (
	"fmt"
	"project1/User"
	"project1/pass"
)

func main() {
	var password string

	Hello := User.SayHello()
	fmt.Println(Hello)

	Name := User.UserName(User.User{})

	fmt.Println("Введите пароль:")
	fmt.Scanln(&password)

	// ok := pass.CheckPassword(password)

	// if !ok {
	// 	fmt.Println("Неверный пароль!")
	// 	return
	// }
	// if ok {
	// 	fmt.Println(ok, "Пароль верный!")
	// }

	if pass.IsStrong(password) {
		fmt.Println("Пароль достаточно сильный!")
	}
	if !pass.IsStrong(password) {
		fmt.Println("Пароль должен быть Не менне 8 символов !")
		return
	}

	fmt.Println("ДОБРО ПОЖАЛОВАТЬ", Name, "!")

}
