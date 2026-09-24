package main

import "fmt"

type Circle struct{ R float64 }
type Rect struct{ W, H float64 }

func circleArea(c Circle) float64 { return 3.14159 * c.R * c.R }
func rectArea(r Rect) float64     { return r.W * r.H }

func main() {
	c := Circle{R: 2}
	r := Rect{W: 4, H: 5}

	fmt.Println("Площадь круга:", circleArea(c))
	fmt.Println("Площадь прямоугольника:", rectArea(r))

	// Проблема: нужен ОДИН срез, где лежат и круги, и прямоугольники,
	// и ОДНА функция, которая посчитает суммарную площадь всех фигур сразу.
	//
	// shapes := []???{c, r}
	//
	// Circle и Rect - совсем разные типы. []Circle не может хранить Rect.
	// Придётся писать totalArea отдельно для кругов и отдельно для прямоугольников,
	// а при добавлении треугольника - переписывать всё заново.
	//
	// Дальше решим эту проблему с помощью интерфейса.
}
