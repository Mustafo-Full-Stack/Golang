package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	ptr := &Point{
		X: 10,
	}
	ptr2 := Point{
		Y: 20,
	}

	fmt.Printf("%+v, %+v", ptr, ptr2)
}
