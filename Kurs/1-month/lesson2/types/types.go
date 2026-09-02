package main

import (
	"fmt"
)



func main() {
	// Три способа объявления
	// var name string = "Алишер" // полная форма
	// var age  = 25               // тип выведется сам
	// city := "Душанбе"          // короткая форма

	//одинарные комменты
	/*

	мы собираемся 
	решать 
	задачи
	*/

	/*fmt.Println(name)
	fmt.Println(age)
	fmt.Println(city)
	*/

	
	//fmt.Println(name, age, city)

	//проверка типов через %T
	//fmt.Printf("%T %T %T\n", name, age, city)

	// //Нулевые значения

	// var a int // 0 
	// var s string // ""
	// var sPoint *string // nil
	// var b bool // false
	// var f float64 // 0.0

	// fmt.Println(a, s, b, f,sPoint)
	// fmt.Printf("%d|%q|%t|%.2f\n", a, s, b, f)

	// // var (
	// // 	d = 0.1
	// // 	e = 0.3 
	// // )
	// d, e := 0.12, 0.39
	// result := d + e 
	// fmt.Printf("result = %f", result)
	//fmt.Println(result)


	// Множественное объявление
	// var x, y int = 10, 20
	// first, second := "раз", "два"
	i, j := 1, 2

	// обмен значениями в одну строку
	i, j = j, i
	fmt.Println(i, j)   // 2 1

	// Ввод от пользователя
	var (
		name string
	 	age int
	)
 
	fmt.Print("Как тебя зовут? ")
	fmt.Scan(&name)
 
	fmt.Print("Сколько тебе лет? ")
	fmt.Scan(&age)
 
	fmt.Printf("%s, тебе %d\n", name, age)
}
