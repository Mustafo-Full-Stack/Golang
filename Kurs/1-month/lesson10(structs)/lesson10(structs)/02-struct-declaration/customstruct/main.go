package customstruct

import "fmt"

type User struct {
	ID      string
	Name    string
	Age     int
	Address string
}

func GetName(u User) string {
	return u.Name
}

func GetProfile(u User) User {
	return u
}

func GetAge(u User) int {
	return u.Age
}

func SetAge(u User) {
	u.Age += 10
	fmt.Println("возраст после SetAgе:", u.Age)
}
