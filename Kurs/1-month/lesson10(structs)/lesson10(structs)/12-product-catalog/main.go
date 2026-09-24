package main

import (
	"fmt"
	"sort"
)

// Урок 10, Шаг 12: мини-проект - каталог товаров.
// собираем всё вместе: структура, срез структур, цикл, форматирование, сортировка

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func main() {
	products := []Product{
		{ID: 1, Name: "Ноутбук", Price: 1200, Quantity: 3},
		{ID: 2, Name: "Мышь", Price: 25.50, Quantity: 10},
		{ID: 3, Name: "Клавиатура", Price: 45, Quantity: 7},
		{ID: 4, Name: "Монитор", Price: 300, Quantity: 2},
	}

	fmt.Printf("%-4s %-14s %10s %5s %12s\n", "ID", "Название", "Цена", "Кол", "Сумма")
	fmt.Println("-----------------------------------------------")

	total := 0.0
	for _, p := range products {
		sum := p.Price * float64(p.Quantity)
		total += sum
		fmt.Printf("%-4d %-14s %10.2f %5d %12.2f\n", p.ID, p.Name, p.Price, p.Quantity, sum)
	}
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Общая стоимость склада: %.2f\n", total) // 4770.00

	// сортировка среза структур по цене
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	fmt.Println("\nСамый дешёвый:", products[0].Name)             // Мышь
	fmt.Println("Самый дорогой:", products[len(products)-1].Name) // Ноутбук
}

// обрати внимание на sort.Slice с анонимной функцией - это урок 9,
// теперь он пригодился на настоящей задаче
