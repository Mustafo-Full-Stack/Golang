package main

import "fmt"

// Урок 10, Шаг 8: вложенная структура - поле, у которого тип тоже структура

type Address struct {
	City   string
	Street string
	Zip    string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	p := Person{
		Name: "Али",
		Age:  25,
		Address: Address{
			City:   "Душанбе",
			Street: "Рудаки 12",
			Zip:    "734000",
		},
	}

	fmt.Println(p.Address.City) // Душанбе - идём по точкам вглубь

	p.Address.Street = "Сомони 5"
	fmt.Printf("%+v\n", p) // {Name:Али Age:25 Address:{City:Душанбе Street:Сомони 5 Zip:734000}}
}

// поле Address имеет имя, поэтому путь всегда полный: p.Address.City.
// на следующем шаге уберём имя поля - и путь станет короче
