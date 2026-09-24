//11.Структура 'Student' ('Name' string,
//'Grades' []int') с методами 'Avarage() float64'
// и 'String() string'.
//Дан срез '[]Student' из 5-6 студентов - с помощью 'Avarage()'
//найди Старосту (Студент с лучшим среднем баллом) и выведи его

package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Avarage() float64 {
	totalNow := 0
	return float64(totalNow) / float64(len(s.Grades))
}

func main() {
	S := []Student{
		{
			Name:   "Samir",
			Grades: []int{90},
		},
		{
			Name:   "Ali",
			Grades: []int{60},
		},
		{
			Name:   "Cat",
			Grades: []int{90},
		},
		{
			Name:   "John",
			Grades: []int{60},
		},
	}

	fmt.Println(
		"Лучший студент:",
		S[0].Name,
	)
	fmt.Println(
		"Средний балл:",
		S[0].Avarage(),
	)
}
