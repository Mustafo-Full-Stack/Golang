package main

import (
	"fmt"

	"myapp/mathutil"
)

func main() {
	fmt.Println(mathutil.Add(3, 4))

	f, _ := mathutil.Factorial(5)
	fmt.Println("5! =", f)

	// Average и MinMax живут в другом файле (stats.go), но импорт
	// не меняется - для main.go пакет mathutil один и тот же.
	avg := mathutil.Average([]float64{4, 8, 15, 16, 23, 42})
	fmt.Printf("Среднее: %.2f\n", avg)

	min, max := mathutil.MinMax([]int{7, 2, 9, -3, 5})
	fmt.Printf("Мин: %d, Макс: %d\n", min, max)
}
