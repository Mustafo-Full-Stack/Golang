package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"` // пустое — не попадёт в JSON
	Discount    float64 `json:"discount,omitempty"`    // 0 — не попадёт в JSON
	CostPrice   float64 `json:"-"`                     // "-" — НИКОГДА не попадёт в JSON
}

func main() {
	p1 := Product{ID: 1, Name: "Ноутбук", Price: 1200, Description: "Мощный", Discount: 10, CostPrice: 800}
	p2 := Product{ID: 2, Name: "Мышь", Price: 25, CostPrice: 10}

	d1, _ := json.MarshalIndent(p1, "", "  ")
	d2, _ := json.MarshalIndent(p2, "", "  ")
	fmt.Println(string(d1))
	fmt.Println(string(d2))

	// У p2 нет description и discount — они пустые.
	// Ни у кого нет cost_price — это закупочная цена, клиенту её видеть нельзя.
	// Так прячут пароли и внутренние данные.
}
