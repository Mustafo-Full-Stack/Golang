package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// switch x := v.(type) работает ТОЛЬКО внутри switch. В каждом case
// переменная x получает конкретный тип этого case.
func describeValue(v any) {
	switch x := v.(type) {
	case nil:
		fmt.Println("Пусто (nil)")
	case int:
		fmt.Printf("Целое %d, удвоенное %d\n", x, x*2)
	case float64:
		fmt.Printf("Дробное %.2f\n", x)
	case string:
		fmt.Printf("Строка %q длиной %d\n", x, len(x))
	case bool:
		fmt.Println("Логическое:", x)
	case []int:
		fmt.Printf("Срез из %d чисел\n", len(x))
	case Shape:
		fmt.Printf("Фигура с площадью %.2f\n", x.Area())
	default:
		fmt.Printf("Неизвестный тип %T\n", x)
	}
}

func main() {
	values := []any{42, "текст", 3.14, true, nil, []int{1, 2}, Circle{Radius: 2}, 'A'}
	for _, v := range values {
		describeValue(v)
	}

	// Обрати внимание на case Shape: сюда попадает любая фигура, у которой есть
	// Area() и Perimeter(). Интерфейсы прекрасно работают и внутри type switch.
}
