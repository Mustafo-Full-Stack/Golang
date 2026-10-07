package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Структура описывает ТОЛЬКО те поля, которые нам нужны.
// Остальные поля из ответа (address, company...) Go просто пропустит.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"surname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Adress   Adress `json:"address"`
}

type GeoCoordinates struct {
	Latitude  string `json:"lat"`
	Longtitue string `json:"lng"`
}

type Adress struct {
	Street  string         `json:"street"`
	Suite   string         `json:"suite"`
	City    string         `json:"city"`
	Zipcode string         `json:"zipcode"`
	Geo     GeoCoordinates `json:"geo"`
}

// Шаг 2. Запрос -> JSON -> структура Go
func main() {
	//url := "https://jsonplaceholder.typicode.com/users/1"
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	// Тот же Unmarshal, что и со строкой в прошлых шагах!
	var u User
	if err := json.Unmarshal(body, &u); err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	fmt.Printf("%+v\n", u)
	fmt.Println("Привет,", u.Name, "! Твоя почта:", u.Email)
}
