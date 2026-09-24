package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {

	result := Car{}

	fmt.Printf("%+v", result)

}
