package mathutil // тот же пакет, что и в mathutil.go!

// Average считает среднее.
func Average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

// MinMax возвращает минимум и максимум. Использует validate
// из mathutil.go - файлы одного пакета видят друг друга без импорта.
func MinMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
