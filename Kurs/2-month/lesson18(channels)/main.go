package main

import (
	"fmt"
	"log"
)

func main() {
	//22 ch := make(chan string)

	// // горутина писатель (значит пишет в канал!)
	// go func() {
	// 	ch <- "4"
	// }()

	// fmt.Println(<-ch)

	/*

		<- ch - мы читаем из канала
		ch <- value - пишем в канал value
	*/

	// chStr := make(chan string)

	// go func() {
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	chStr <- "hello1"
	// 	chStr <- "hello2"
	// 	chStr <- "hello2"
	// 	close(chStr)
	// }()

	// for msg := range chStr {
	// 	fmt.Println(msg)
	// }

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	ch <- 3
	_, closed := <-ch
	if closed {
		log.Println("канал закрыть")
	}
	// for value := range ch {
	// 	log.Println(value)
	// }
	//fmt.Println(val)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
