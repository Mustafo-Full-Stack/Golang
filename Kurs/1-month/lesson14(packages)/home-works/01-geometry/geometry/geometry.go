// Задание: пакет geometry.
// Реализуй экспортируемые функции:
//
//	RectangleArea(width, height float64) float64
//	RectanglePerimeter(width, height float64) float64
//
// Функции должны вычислять площадь и периметр прямоугольника.
package geometry

func RectangleArea(width, height float64) float64 {

	return width * height
}

func RectanglePerimeter(width, height float64) float64 {

	return width + height + width + height
}
