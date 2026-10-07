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
	u := User{
		Name:  "Иван",
		Age:   16,
		Email: "mustafodropshiping4@gmail.com",
	}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(data))
}
