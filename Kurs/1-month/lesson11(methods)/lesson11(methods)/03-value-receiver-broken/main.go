package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) ScaleBroken(factor float64) {
	r.Width *= factor
	r.Height *= factor
	fmt.Printf("Внутри метода: %.1fx%.1f\n", r.Width, r.Height)
}

func (r *Rectangle) ScaleBrokenWithPtr(factor float64) {
	r.Width *= factor
	r.Height *= factor
	fmt.Printf("Внутри метода ScaleBrokenWithPtr: %.1fx%.1f\n", r.Width, r.Height)
}

func main() {
	r := Rectangle{Width: 3, Height: 4}

	r.ScaleBrokenWithPtr(2)
	fmt.Printf("Снаружи: %.1fx%.1f\n", r.Width, r.Height)

	// r.ScaleBroken(2)
	// fmt.Printf("Снаружи: %.1fx%.1f\n", r.Width, r.Height)
}
