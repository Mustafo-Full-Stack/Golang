package main

import "fmt"

func main() {
	var slice = make([]int, 5)
	for i := 0; i < 5; i++ {
		slice = append(slice, 1)
		fmt.Println(slice, "Длинна:", len(slice), "Ёмкость", cap(slice))
	}
}
