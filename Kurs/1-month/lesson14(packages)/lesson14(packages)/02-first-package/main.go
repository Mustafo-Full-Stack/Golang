package main

import (
	"fmt"

	"myapp/mathutil"
)

func main() {
	fmt.Println(mathutil.Add(3, 4))
	fmt.Println(mathutil.Max(10, 7))

	f, err := mathutil.Factorial(5)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("5! =", f)

	_, err = mathutil.Factorial(-3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	//mathutil.validate(5)
}
