package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

func (r Rectangle) Describe() string {
	return fmt.Sprintf("Прямоугольник %.1fx%.1f, площадь %.2f",
		r.Width, r.Height, r.Area())
}

func main() {
	r := Rectangle{Width: 3, Height: 4}


	fmt.Println(r.Area(), r.Perimeter(), r.IsSquare())
	fmt.Println(r.Describe())

	sq := Rectangle{Width: 5, Height: 5}
	fmt.Println(sq.IsSquare())
}
