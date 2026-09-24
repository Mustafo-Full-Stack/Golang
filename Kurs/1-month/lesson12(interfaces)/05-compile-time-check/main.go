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

// Эта строка ничего не делает во время выполнения программы - переменная
// нигде не используется. Зато она заставляет компилятор ПРЯМО СЕЙЧАС
// проверить: реализует ли Circle интерфейс Shape. Если нет - сборка упадёт
// с понятной ошибкой, а не где-то потом в рантайме.
var _ Shape = Circle{}
var _ Shape = Rectangle{}

func main() {
	fmt.Println("Если программа собралась - обе фигуры точно реализуют Shape")

	// Убери у Circle метод Perimeter и посмотри на ошибку компилятора:
	// ./main.go:XX:14: cannot use Circle{} (value of type Circle) as Shape value
	// in variable declaration: Circle does not implement Shape (missing method Perimeter)
}
