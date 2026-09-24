package main

import "fmt"

type Animal interface {
	Sound()
}

type Cat struct {
}

func (c Cat) Sound() {
	fmt.Println("Мяу")
}

func MakeSound(a Animal) {
	a.Sound()
}

func main() {
	MakeSound(Cat{})
}
