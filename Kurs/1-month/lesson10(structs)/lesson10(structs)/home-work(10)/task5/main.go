package main

import "fmt"

type Laptop struct {
	Brand string
	RAM   int
	SSD   int
	Price float64
}

func main() {
	info := Laptop{
		Brand: "HP ENVY",
		RAM:   8,
		SSD:   512,
		Price: 780,
	}

	fmt.Println("BRAND :", info.Brand)
	fmt.Println("RAM :", info.RAM)
	fmt.Println("SSD :", info.SSD)
	fmt.Println("PRICE :", info.Price, "$")

}
