package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Introduce() {
	fmt.Println("Меня зовут", u.Name)
	fmt.Println("Мне", u.Age, "лет")
}

func main() {
	u := User{
		Name: "Alex",
		Age:  20,
	}

	u.Introduce()
}
