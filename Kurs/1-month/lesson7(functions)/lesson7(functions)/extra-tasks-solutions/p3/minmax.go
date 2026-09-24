package main

import "fmt"

// функция возвращает сразу два значения - минимум и максимум среза
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	nums := []int{4, 8, 1, 9, 3, 6}

	min, max := minMax(nums)
	fmt.Println("min:", min, "max:", max)
}
