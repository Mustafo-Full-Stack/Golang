package main

import "fmt"

func main() {
	var slice = []int{10, 20, 30, 40}
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Println(slice[i])
	}
}
