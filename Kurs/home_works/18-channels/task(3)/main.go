package main

import "fmt"

func main() {
	ch := make(chan int)

	//ch <- 10
	//deadlock, чтение никогда не произойдет и
	//потом программа зависает навсегда :)

	fmt.Println(ch)
}
