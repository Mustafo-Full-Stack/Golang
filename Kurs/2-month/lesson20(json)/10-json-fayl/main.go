package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// v any — подойдёт любой тип (any вы уже видели в уроке про интерфейсы)
func saveJSON(filename string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("упаковка в JSON: %w", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("запись файла: %w", err)
	}
	return nil
}

func loadJSON(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("чтение файла: %w", err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("разбор JSON: %w", err)
	}
	return nil
}

func main() {
	users := []User{
		{Name: "Али", Age: 25, Email: "ali@mail.tj"},
		{Name: "Вера", Age: 30, Email: "vera@mail.tj"},
	}

	if err := saveJSON("users.json", users); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Сохранено в users.json")

	var loaded []User
	if err := loadJSON("users.json", &loaded); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Загружено %d: %+v\n", len(loaded), loaded)
}
