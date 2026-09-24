package main

import "fmt"

// Урок 10, Шаг 3: четыре способа создать структуру

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	u1 := User{Name: "Али", Age: 25, Email: "ali@mail.tj"} // с именами полей - так и надо
	u2 := User{"Вера", 30, "vera@mail.tj"}                 // по порядку - хрупко
	var u3 User                                            // все поля нулевые
	u4 := User{Name: "Тимур"}                              // только часть полей

	fmt.Printf("%+v\n", u1) // {Name:Али Age:25 Email:ali@mail.tj}
	fmt.Printf("%+v\n", u2) // {Name:Вера Age:30 Email:vera@mail.tj}
	fmt.Printf("%+v\n", u3) // {Name: Age:0 Email:}
	fmt.Printf("%+v\n", u4) // {Name:Тимур Age:0 Email:}
}

// обрати внимание на u3: нулевые значения из урока 2 работают и здесь.
//
// покажи вживую, почему второй способ опасен: добавь в User поле
// Phone string между Name и Age - и строка с u2 перестанет компилироваться:
// cannot use 30 (untyped int constant) as string value in struct literal
// too few values in struct literal of type User
//
// а строки с именами полей продолжат работать как ни в чём не бывало.
// и учти: будь новое поле тоже string, код собрался бы молча,
// но значения легли бы не в те поля - это хуже ошибки компиляции
