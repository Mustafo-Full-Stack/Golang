package main

import "fmt"

// Урок 10, Шаг 9: встраивание - поле без имени, его поля "поднимаются" наверх

type Address struct {
	City   string
	Street string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Employee struct {
	Person   // встроено: имени поля нет, есть только тип
	Position string
	Salary   float64
}

func main() {
	e := Employee{
		Person: Person{
			Name:    "Вера",
			Age:     30,
			Address: Address{City: "Худжанд"},
		},
		Position: "Разработчик",
		Salary:   5000,
	}

	// fmt.Println(e.Name)         // Вера - напрямую, не e.Person.Name
	// fmt.Println(e.Address.City) // Худжанд
	// fmt.Println(e.Position)     // Разработчик

	// полная форма тоже доступна
	// fmt.Println(e.Person.Name) // Вера

	fmt.Printf("%+v\n", e)

	e.Person.Age = 31
	fmt.Println(e.Person.Age)
}

// сравни с шагом 8: там Address было полем с именем, здесь Person - без имени.
// это основа композиции в Go, на уроке 11 из неё вырастет
// "композиция вместо наследования"
