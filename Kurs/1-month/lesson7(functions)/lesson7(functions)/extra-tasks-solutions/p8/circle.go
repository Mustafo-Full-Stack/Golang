package main

import (
	"fmt"
	"math"
)

// именованные результаты area и perimeter уже объявлены
// в скобках после параметра - сразу видно, что вернёт функция
func circleAreaPerimeter(r float64) (area, perimeter float64) {
	area = math.Pi * r * r
	perimeter = 2 * math.Pi * r
	return
}

func main() {
	var r float64
	fmt.Print("Введите радиус r:")
	fmt.Scanln(&r)

	area, perimeter := circleAreaPerimeter(r)
	fmt.Printf("Площадь: %.2f, периметр: %.2f\n", area, perimeter)
}
