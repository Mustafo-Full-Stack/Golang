package main

import "fmt"

func main() {
	var nums []int

	fmt.Println("Вводи числа (0 — закончить):")
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		nums = append(nums, n)
	}

	fmt.Println("\nИсходный срез:", nums)

	if len(nums) == 0 {
		fmt.Println("Ничего не ввели")
		return
	}

	for i := 0; i < len(nums); i++ {
		// Проверяем: видели ли мы это число раньше в массиве?
		alreadyPrinted := false
		for k := 0; k < i; k++ { // только до текущей позиции i
			if nums[k] == nums[i] {
				alreadyPrinted = true
				break
			}
		}

		if alreadyPrinted {
			continue // переходим к следующей итерации
		}

		// Считаем сколько раз встречается nums[i]
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}

		fmt.Printf("Число %d встречается %d раз(а)\n", nums[i], count)
	}
}
