package main

import "fmt"

func calculate(nums []int, result chan<- int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	result <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int)

	go calculate(nums[:5], result)
	go calculate(nums[5:], result)

	sum1 := <-result
	sum2 := <-result
	fmt.Println("Общая сумма:", sum1+sum2) // 55

	// Красиво: две горутины считают половинки, main собирает результат.
	// Никаких мьютексов не понадобилось!
}
