package main

import "fmt"

/*
Форма 1 — классическая (со счётчиком)
for инициализация(начальное условие); условие; шаг {
	// тело
}

*/

func main() {
	// i ++ -> i = i + 1
	// for i := 1; i <= 5; i++ {
	// 	// i = 1 1<=5 (true)
	// 	// i = 2 2 <=5 (true)
	// 	// i =5 5 <=5 (true)
	// 	// i = 6 (break)
	// 	fmt.Println("Итерация номер", i)
	// }
	// fmt.Println("Цикл закончился")

	// что выведет?
	// for i := 0; i < 0; i++ {
	// fmt.Println("hi")
	// }

	// i = i + 2
	// for i := 0; i <= 10; i += 2 {
	// 	fmt.Print(i, " ") // 0 2 4 6 8 10
	// }
	// fmt.Println()

	// обратный отсчёт
	// for i := 10; i > 0; i-- {
	// 	fmt.Print(i, " ") // 10 9 8 ... 1
	// }
	// fmt.Println("Пуск!")

	//cумма чисел
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	// i = 1
	// 	// sum = 0
	// 	// sum = 0 + 1
	// 	//sum = 1
	// 	sum = sum + i
	// 	//sum += i
	// }
	// fmt.Printf("Сумма чисел от 1 до 100:%d", sum)

	//Сумма всех нечётных чисел от 1 до 1000
	sum := 0
	for i := 1; i <= 1000; i++ {
		if i%2 != 0 {
			//sum += i
			sum = sum + i
			//сокращенная запись sum = sum + i
		}
	}
	fmt.Printf("Сумма:%d", sum)

	//вывод чисел от 100 до 90 в обратном порядке.
	// for i := 100; i >= 90; i-- {
	// 	fmt.Printf("%d ", i)
	// }

	//Вывести числа от 1 до 100,
	// но вместо кратных 3 писать «Fizz», кратных 5 — «Buzz»,
	// кратных и 3 и 5 — «FizzBuzz».
	// for i := 1; i <= 100; i++ {
	// 	switch {
	// 	case i%3 == 0 && i%5 == 0:
	// 		fmt.Println("FizzBuzz")
	// 	case i%3 == 0:
	// 		fmt.Println("Fizz")
	// 	case i%5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }
}
