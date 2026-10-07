package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Когда структуру заранее не знаем — разбираем в map[string]any
	raw := `{"name":"Али","age":25,"active":true,"scores":[90,85],"meta":{"role":"admin"}}`

	var data map[string]any
	json.Unmarshal([]byte(raw), &data)

	for k, v := range data {
		fmt.Printf("%s = %v  (тип %T)\n", k, v, v)
	}

	// ВАЖНО: age — это float64, а не int! В JSON все числа одного вида.
	// Достаём значение безопасно: v, ok := x.(тип)
	if age, ok := data["age"].(float64); ok {
		fmt.Println("Возраст:", int(age))
	}
	if meta, ok := data["meta"].(map[string]any); ok {
		fmt.Println("Роль:", meta["role"])
	}

	// Правило: знаем структуру — используем struct. Не знаем — map[string]any.
}
