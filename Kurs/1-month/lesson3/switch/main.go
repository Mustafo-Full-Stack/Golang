package main

import "fmt"

func main() {
	// day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Понедельник")
	// case 2:
	// 	fmt.Println("Вторник")
	// case 3:
	// 	fmt.Println("Среда")
	// case 6, 7: // несколько значений в одном case
	// 	fmt.Println("Выходной!")
	// default:
	// 	fmt.Println("Другой день")
	// }

	// score := 85
	// //switch без выражения (замена длинных if)
	// switch {
	// case score >= 90:
	// 	fmt.Println("Отлично")
	// case score >= 75:
	// 	fmt.Println("Хорошо")
	// case score >= 60:
	// 	fmt.Println("Удовлетворительно")
	// default:
	// 	fmt.Println("Неудовлетворительно")
	// }

	// a := 3
	// switch a {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// 	break
	// default:
	// 	fmt.Println("default")
	// }
	
	// fallthrough — это ключевое слово в Go, 
	// которое заставляет выполнить 
	// следующий case, даже если его условие не подходит.
	a := 2

	switch a {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
		fallthrough

	case 3:
		fmt.Println("three")
		fallthrough

	default:
		fmt.Println("default")
	}
}
