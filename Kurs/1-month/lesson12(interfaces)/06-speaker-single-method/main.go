package main

import "fmt"

// Чем меньше интерфейс, тем он полезнее. Один метод - идиома Go.
type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }
type Robot struct{ ID int }

func (d Dog) Speak() string   { return d.Name + ": Гав!" }
func (c Cat) Speak() string   { return c.Name + ": Мяу!" }
func (r Robot) Speak() string { return fmt.Sprintf("Робот-%d: Бип-буп", r.ID) }

func main() {
	speakers := []Speaker{
		Dog{Name: "Шарик"},
		Cat{Name: "Мурка"},
		Robot{ID: 42},
	}

	for _, s := range speakers {
		fmt.Println(s.Speak())
	}

	// Робот - не животное, у него нет ничего общего с Dog и Cat по смыслу.
	// Но интерфейсу Speaker всё равно, "кто ты" - важно только умение Speak().
	// В Go на первом месте способность, а не происхождение или родство типов.
}
