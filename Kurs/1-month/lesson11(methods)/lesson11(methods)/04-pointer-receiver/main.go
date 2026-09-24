package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	fmt.Printf("После Scale: %.1fx%.1f\n", r.Width, r.Height)
}
