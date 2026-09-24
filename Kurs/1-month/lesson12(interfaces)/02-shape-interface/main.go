package main

import (
	"fmt"
	"math"
)

// Shape - это не сущность, а "умение". Любой тип с методами Area и Perimeter
// автоматически подходит под этот интерфейс. Писать "Circle implements Shape" не нужно.
type Shape interface {
	Area() float64
	Perimeter() float64
	IsSquare() bool
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (r Circle) IsSquare() bool {
	return false
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (r Rectangle) IsSquare() bool {
	return r.Height == r.Width
}

func main() {
	var s Shape

	// s = Circle{Radius: 5}
	// fmt.Printf("Круг: площадь %.2f\n", s.Area())

	s = Rectangle{Width: 4, Height: 4}
	fmt.Printf("Прямоугольник: площадь %.2f\n", s.Area())
	fmt.Println(s.IsSquare())

	// Нигде не написано "Circle implements Shape" и "Rectangle implements Shape".
	// Компилятор сам увидел, что у типов есть нужные методы, и разрешил присвоить
	// их переменной типа Shape. Это и есть неявная реализация.
}
