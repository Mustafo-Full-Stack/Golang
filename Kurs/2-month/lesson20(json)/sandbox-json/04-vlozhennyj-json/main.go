package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Geo    Geo    `json:"geo"` // объект внутри объекта внутри объекта
}

type Company struct {
	Name string `json:"name"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
	Company Company `json:"company"`
}

// Шаг 4. Вложенный JSON = вложенные структуры.
// Рисуем структуры "матрёшкой" так же, как устроен JSON.
func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	for _, u := range users {
		fmt.Printf("%-25s | %-15s | %s | %s\n",
			u.Name, u.Address.City, u.Company.Name, u.Address.Geo.Lat)
	}
}
