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

// Triangle появился только что. Ни Shape, ни describe, ни totalArea не менялись.
type Triangle struct {
	A, B, C float64
}

func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

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
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Triangle{A: 3, B: 4, C: 5}, // просто добавили новый тип
	}

	for _, s := range shapes {
		describe(s)
	}

	// describe и totalArea не изменились ни на строку. Это и есть расширяемость:
	// новый тип фигуры добавляется без единой правки старого кода.
	fmt.Printf("Суммарная площадь: %.2f\n", totalArea(shapes))
}
