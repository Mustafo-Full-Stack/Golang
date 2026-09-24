package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// describe и totalArea ничего не знают ни про Circle, ни про Rectangle.
// Им достаточно того, что перед ними Shape - что-то умеющее Area() и Perimeter().
func describe(s Shape) {
	fmt.Printf("Площадь: %8.2f | Периметр: %8.2f\n", s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	// Вот и решение проблемы из шага 1: разные типы лежат в одном срезе,
	// потому что все они - Shape.
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 1},
	}

	for _, s := range shapes {
		describe(s)
	}

	fmt.Printf("Суммарная площадь: %.2f\n", totalArea(shapes))
}
