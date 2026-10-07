package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	inputs := []string{
		`{"name":"Али","age":25}`,         // всё хорошо
		`{"name":"Али","age":"двадцать"}`, // не тот тип
		`{"name":"Али",}`,                 // синтаксическая ошибка
		``,                                // пустая строка
	}

	for i, in := range inputs {
		var u User
		err := json.Unmarshal([]byte(in), &u)
		if err != nil {
			fmt.Printf("%d. Ошибка: %v\n", i+1, err)
			continue
		}
		fmt.Printf("%d. OK: %+v\n", i+1, u)
	}

	// Данным извне не доверяем: всегда проверяем err после Unmarshal.
}
