package main

import (
	"fmt"

	"geometry-homework/geometry"
)

// Домашнее задание 1: покажи площадь и периметр 2–3 прямоугольников.
func main() {
	width1, height1 := 4.0, 4.0
	fmt.Println("4 x 4:",
		"площадь =", geometry.RectangleArea(width1, height1),
		"периметр =", geometry.RectanglePerimeter(width1, height1))

	// TODO: добавить ещё один или два прямоугольника с разными сторонами.
}
