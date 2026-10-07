package main

import (
	"encoding/json"
	"fmt"
)

// Ошибка 1: поля с маленькой буквы
type badUser struct {
	name string
	age  int
}

// Ошибка 3: пробел в теге (после двоеточия)
type spaceTag struct {
	Name string `json: "name"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 1. Маленькие буквы -> пустой JSON, без ошибки
	data, _ := json.Marshal(badUser{"Али", 25})
	fmt.Println("1.", string(data)) // {}

	// 2. Забыли & -> ошибка
	var u User
	err := json.Unmarshal([]byte(`{"name":"Али"}`), u)
	fmt.Println("2.", err)

	// 3. Пробел в теге -> тег не работает, ошибки нет
	data, _ = json.Marshal(spaceTag{"Али"})
	fmt.Println("3.", string(data)) // {"Name":"Али"}, а не {"name":"Али"}

	// 4. Неверный тип
	err = json.Unmarshal([]byte(`{"age":"25"}`), &u)
	fmt.Println("4.", err)

	// 5. Число из map[string]any — float64, а не int
	var m map[string]any
	json.Unmarshal([]byte(`{"n":5}`), &m)
	fmt.Printf("5. тип n = %T\n", m["n"])
	// n := m["n"].(int) // <- раскомментируйте: panic: ... is float64, not int
}
