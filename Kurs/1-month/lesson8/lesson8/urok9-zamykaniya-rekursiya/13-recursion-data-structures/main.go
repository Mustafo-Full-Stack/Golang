package main

import "fmt"

// Урок 9, Шаг 11: рекурсия на структурах данных

func sumSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sumSlice(nums[1:])
}

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}

func main() {
	fmt.Println(sumSlice([]int{1, 2, 3, 4, 5})) // 15
	fmt.Println(reverseString("golang"))        // gnalog
}
