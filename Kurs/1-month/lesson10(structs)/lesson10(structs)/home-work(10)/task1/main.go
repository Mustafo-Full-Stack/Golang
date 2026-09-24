package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {
	u1 := Book{
		Title:  "Гулистон",
		Author: "Саъди Шерози",
		Year:   1921,
		Price:  500.00,
	}

	fmt.Printf("%+v", u1)
}
