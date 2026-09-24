package User

import "fmt"

type User struct {
	Name string
}

func UserName(u User) string {
	fmt.Println("Введите Имя:")
	fmt.Scanln(&u.Name)
	return u.Name
}

func SayHello() string {
	return "Привет из пакета user!"
}
