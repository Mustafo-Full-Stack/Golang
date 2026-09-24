package main

import "fmt"

// Урок 10, Шаг 11: анонимная структура и сравнение структур

type User struct {
	Name string
	Age  int
}

func main() {
	// одноразовая структура - отдельный тип объявлять не стали
	config := struct {
		Host  string
		Port  int
		Debug bool
	}{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
	}
	fmt.Printf("%+v\n", config) // {Host:localhost Port:8080 Debug:true}

	// структуры сравниваются через ==, если все их поля сравнимы
	a := User{Name: "Али", Age: 25}
	b := User{Name: "Али", Age: 25}
	fmt.Println(a == b) // true - сравниваются поля, а не адреса

	b.Age = 26
	fmt.Println(a == b) // false
}

// оговорка: если внутри структуры есть срез или map, == работать не будет.
// это увидим в папке 13-tipichnye-oshibki
