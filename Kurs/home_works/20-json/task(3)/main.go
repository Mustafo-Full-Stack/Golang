package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

func main() {
	p := Product{
		Name:    "Шоколад",
		Price:   20,
		ID:      1,
		InStock: true,
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(data))
}
