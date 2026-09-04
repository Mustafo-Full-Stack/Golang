package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	//// удалить элемент с индексом 2 ("c")
	//s = append(s[:2], s[3:]...)
	nums = append(nums[2:4])
	fmt.Println(nums, len(nums), cap(nums))
}
