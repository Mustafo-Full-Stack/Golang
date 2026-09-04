package main

import "fmt"

func main() {
	// zeros := make([]int, 5, 15)
	// // []int
	// fmt.Println(zeros)
	// fmt.Println("len=",len(zeros),"cap=",cap(zeros))

	// длина 0, но место под 100 — без перевыделений
	nums := make([]int, 0, 100)
	//fmt.Println(len(nums), cap(nums))

	for i := 0; i < 101; i++ {
		nums = append(nums, i) // ни одного перевыделения
	}

	fmt.Println(len(nums), cap(nums)) 
}
