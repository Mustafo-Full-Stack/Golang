package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Привет из горутины!"
	}()

	v := <-ch
	fmt.Println(v)
}
