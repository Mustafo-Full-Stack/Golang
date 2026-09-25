package main

import (
	"fmt"
	"project2/users"
)

func main() {
	// Тип User импортирован из отдельного пакета users.
	user1 := users.User{Name: "Али", Age: 17}
	user2 := users.User{Name: "Вали", Age: 21}
	user3 := users.User{Name: "Сайвали", Age: 18}

	// Метод меняет возраст исходного пользователя через указатель.
	user1.HaveBirthday()
	user2.HaveBirthday()
	user3.HaveBirthday()

	fmt.Printf("%+v, %+v\n\n", user1, user1.IsAdult())
	fmt.Printf("%+v, %+v\n\n", user2, user2.IsAdult())
	fmt.Printf("%+v, %+v\n\n", user3, user3.IsAdult())
}

/*
Предыдущие упражнения:

alif.Println(User.SayHello("Али"))
alif.Println(person.IsAge())
alif.Println(person2.IsAge())
alif.Println(User.IsEven(20))
*/
