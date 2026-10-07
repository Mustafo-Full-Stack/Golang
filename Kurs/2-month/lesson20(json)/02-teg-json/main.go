package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Тег — подсказка для пакета json: как назвать поле в JSON
// Пишется в обратных кавычках. Пробела после двоеточия НЕТ!
type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func main() {
	u := User{
		Name:    "Али",
		Surname: "Алиев",
		Age:     25,
		Email:   "ali@mail.tj",
		IsAdmin: false,
	}

	data, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(data))

	// В Go поле называется Name, а в JSON — name. Тег — это "переводчик".
}
