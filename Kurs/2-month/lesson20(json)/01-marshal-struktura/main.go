package main

import (
	"encoding/json"
	"fmt"
)

// Поля — с БОЛЬШОЙ буквы: пакет json видит только их
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	u := User{Name: "Али", Age: 25, Email: "ali@mail.tj"}

	// Marshal — "упаковать чемодан": структура Go -> JSON
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Marshal возвращает []byte, поэтому превращаем в строку
	fmt.Println(string(data))
	// {"Name":"Али","Age":25,"Email":"ali@mail.tj"}

	// Вопрос группе: в JSON принято писать name с маленькой буквы. Как быть?
	// Ответ — в следующем шаге (теги).
}
