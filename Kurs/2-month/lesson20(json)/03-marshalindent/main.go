package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := User{Name: "Али", Age: 25, Email: "ali@mail.tj"}

	// MarshalIndent — то же самое, но с отступами, чтобы человек мог прочитать
	// "" — префикс строки, "  " — один уровень отступа (2 пробела)
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))

	// Срез структур превращается в JSON-массив
	users := []User{
		{Name: "Али", Age: 25, Email: "ali@mail.tj"},
		{Name: "Вера", Age: 30, Email: "vera@mail.tj"},
	}
	data, _ = json.MarshalIndent(users, "", "  ")
	fmt.Println(string(data))
}
