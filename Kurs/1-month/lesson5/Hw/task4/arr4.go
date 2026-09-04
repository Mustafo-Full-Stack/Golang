package main

import "fmt"

func main() {
	nums := []string{}
	nums = append(nums, "Hello", "World", "dd!")
	fmt.Println(nums[0], len(nums[0]), cap(nums))
	fmt.Println(nums[1], len(nums[1]), cap(nums))
	fmt.Println(nums[2], len(nums[2]), cap(nums))
}
