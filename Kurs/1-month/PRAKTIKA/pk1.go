package main

import "fmt"

type Car interface {
	Info()
}

type Brand struct {
}

func (b Brand) Info() {
	fmt.Println("Инфo: BMW 2026 X5")
}

func ShowInfo(c Car) {
	c.Info()
}

func main() {
	ShowInfo(Brand{})
}
