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
	// В JSON больше полей, чем в структуре — лишние игнорируются
	extra := `{"name":"Али","age":25,"email":"a@b.c","city":"Душанбе","phone":"123"}`
	var u1 User
	json.Unmarshal([]byte(extra), &u1)
	fmt.Printf("%+v\n", u1) // city и phone просто отброшены

	// В JSON меньше полей — остальные получат нулевое значение
	partial := `{"name":"Тимур"}`
	var u2 User
	json.Unmarshal([]byte(partial), &u2)
	fmt.Printf("%+v\n", u2) // {Name:Тимур Age:0 Email:}

	// Go не ругается ни на лишнее, ни на недостающее.
	// Плюс: API можно расширять, старый код не ломается.
	// Минус: опечатка в теге молча даёт пустое поле.
}
