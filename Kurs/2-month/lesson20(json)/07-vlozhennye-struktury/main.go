package main

import (
	"encoding/json"
	"fmt"
)

// Объект внутри объекта = структура внутри структуры
type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Employee struct {
	Name      string   `json:"name"`
	Salary    float64  `json:"salary"`
	Skills    []string `json:"skills"`     // массив JSON = срез Go
	Address   Address  `json:"address"`    // объект JSON = структура Go
	ManagerID *int     `json:"manager_id"` // указатель: может быть null
}

func main() {
	// ---- Go -> JSON ----
	managerID := 5
	emps := []Employee{
		{Name: "Али", Salary: 5000, Skills: []string{"Go", "SQL"},
			Address: Address{City: "Душанбе", Street: "Рудаки 12"}, ManagerID: &managerID},
		{Name: "Вера", Salary: 7000, Skills: []string{"Go"},
			Address: Address{City: "Худжанд", Street: "Ленина 5"}, ManagerID: nil},
	}
	data, _ := json.MarshalIndent(emps, "", "  ")
	fmt.Println(string(data))
	// У Веры manager_id: null — nil-указатель стал null

	// ---- JSON -> Go ----
	input := `[
	  {"name":"Тимур","salary":4000,"skills":["Python"],
	   "address":{"city":"Бохтар","street":"Саид 3"},"manager_id":null}
	]`
	var loaded []Employee
	if err := json.Unmarshal([]byte(input), &loaded); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	for _, e := range loaded {
		fmt.Println(e.Name, e.Address.City, e.Skills, e.ManagerID == nil)
	}
}
