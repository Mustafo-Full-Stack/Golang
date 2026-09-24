package main

import "fmt"

// Урок 10, Шаг 4: доступ к полям через точку и изменение полей

type User struct {
	Name  string
	Age   int
	Email string
	LastName string
}

func main() {
	u1 := User{Name: "Али", Age: 25, Email: "ali@mail.tj"}

	fmt.Println(u1.Name) // Али

	u1.Age = 26
	u1.Email = "new@mail.tj"
	fmt.Printf("%+v\n", u1) // {Name:Али Age:26 Email:new@mail.tj}

	// поле - это обычное значение, его можно использовать в выражениях
	if u1.Age >= 18 {
		fmt.Println(u1.Name, "- совершеннолетний") // Али - совершеннолетний
	}
}
