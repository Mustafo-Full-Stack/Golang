package main

import "fmt"

type Studet struct {
	Name   string
	Grade  int
	Passed bool
}

func main() {
	result := Studet{
		Name: "Мистер Бист",
	}

	fmt.Printf("%+v", result)
}
