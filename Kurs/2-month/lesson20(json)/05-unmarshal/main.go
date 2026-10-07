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
	// JSON обычно приходит откуда-то снаружи: из файла, из интернета
	input := `{"name":"Вера2","age":30,"email":"vera@mail.tj"}`

	// Unmarshal — "распаковать чемодан": JSON -> структура Go
	var u User
	err := json.Unmarshal([]byte(input), &u) // &u — Unmarshal должен ИЗМЕНИТЬ u
	if err != nil {
		fmt.Println("Ошибка разбора:", err)
		return
	}

	fmt.Printf("%+v\n", u)
	fmt.Println("Имя:", u.Name)
	fmt.Println("Через год будет:", u.Age+1)

	// Почему &u? Помните урок про указатели: без & функция получит КОПИЮ
	// и наша переменная u останется пустой.
}
