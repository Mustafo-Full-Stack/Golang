package main

import "fmt"

type Speaker interface {
	Speak() string
}

// Counter специально с pointer receiver - для разбора поломки 2.
type Counter struct{ n int }

func (c *Counter) Speak() string { return "тик" }

func main() {
	// Правильно: передаём указатель, а не значение.
	var s Speaker = &Counter{}
	fmt.Println(s.Speak())

	// поломка 2 - самая частая ошибка новичков с интерфейсами.
	// var s2 Speaker = Counter{}
	// Counter does not implement Speaker (method Speak has pointer receiver)
	//
	// Правило: если метод объявлен с pointer receiver (func (c *Counter) ...),
	// то интерфейсу удовлетворяет ТОЛЬКО указатель *Counter, а не само значение Counter.
	// Решение - писать &Counter{} вместо Counter{}.

	// поломка 3 - паника при неверном type assertion.
	var i any = 42
	// str := i.(string)
	// panic: interface conversion: interface {} is int, not string
	_ = i

	// поломка 4 - .(type) работает только внутри switch.
	// t := i.(type)
	// use of .(type) outside type switch

	// поломка 5 - nil-интерфейс: переменная объявлена, но ничем не заполнена.
	var empty Speaker
	fmt.Println(empty == nil) // true - внутри пусто, ни типа, ни значения
	// fmt.Println(empty.Speak())
	// panic: runtime error: invalid memory address or nil pointer dereference

	// поломка 1 - не хватает метода в реализации (разбирается в 05-compile-time-check):
	// var _ Shape = Square{} // missing method Perimeter
	// Читайте сообщение компилятора вместе с группой - оно прямо называет
	// недостающий метод.
}
