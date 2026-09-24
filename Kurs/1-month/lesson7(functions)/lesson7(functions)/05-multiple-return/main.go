package main

import "fmt"

// Go умеет возвращать сразу несколько значений - другие языки обычно
// возвращают только одно. Это фундамент философии ошибок в Go:
// значение, ошибка := функция() - увидим позже, на уроке про ошибки

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

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
	q, r := divmod(17, 5)
	fmt.Printf("17 / 5 = %d остаток %d\n", q, r)

	lo, hi := minMax([]int{4, 8, 1, 9, 3})
	fmt.Println("min:", lo, "max:", hi)

	// если один из результатов не нужен - игнорируем его через _
	_, remainder := divmod(17, 5)
	fmt.Println("Только остаток:", remainder)

	// поломка: нельзя принять два возвращаемых значения в одну переменную
	// q2 := divmod(17, 5)
	// ошибка: assignment mismatch: 1 variable but divmod returns 2 values
}
