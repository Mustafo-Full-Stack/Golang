package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Положили 3 значения, никто не читал - и ничего не заблокировалось")
	fmt.Println("len =", len(ch), "cap =", cap(ch)) // 3 3

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println("len =", len(ch)) // 1

	// len - сколько сейчас в буфере, cap - вместимость. Как со срезами!

	// Переполнение (раскомментируй, чтобы увидеть deadlock):
	// ch <- 4
	// ch <- 5
	// ch <- 6 // deadlock - буфер полон
}
