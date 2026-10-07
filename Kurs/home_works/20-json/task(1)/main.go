package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

// func (b Book) Bookstruct() Book {
// }

func main() {
	b := Book{
		Title:  "Евгений Онегин",
		Author: "А. С. Пушкин",
		Year:   2010,
		Price:  500,
	}

	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(data))

}
