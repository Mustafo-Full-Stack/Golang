// package main

// import "fmt"

// // создаем свой тип для красоты кода
// type Length float64

// type Rectangle struct {
// 	Width  Length
// 	Height Length
// }

// // обычная функция - работает "снаружи" типа
// func area(r Rectangle) float64 {
// 	return float64(r.Height * r.Width)
// }

// // Area returns the area of the rectangle.
// // метод - та же логика, но "прикреплена" к Rectangle
// // через receiver (r Rectangle)
// func (r Rectangle) Area() float64 {
// 	return float64(r.Width * r.Height)
// }

// func main() {
// 	rect := Rectangle{
// 		Width:  Length(10),
// 		Height: Length(3),
// 	}

// 	fmt.Println(area(rect))

// 	fmt.Println(rect.Area())
// 	//rect.Area()
// }

package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) String() string {
	return s.Name
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	total := 0
	for _, nums := range s.Grades {
		total += nums
	}

	avg := float64(total) / float64(len(s.Grades))
	return avg
}

func main() {
	students := []Student{
		{
			Name:   "Ali",
			Grades: []int{1, 2, 3, 5, 10},
		},
		{
			Name:   "Akbar",
			Grades: []int{1, 2, 10, 5, 8},
		},
		{
			Name:   "Anush",
			Grades: []int{1, 2, 3, 0, 2},
		},
		{
			Name:   "Shahrom",
			Grades: []int{1, 2, 9, 7, 7},
		},
	}

	bestAvgGradeStd := students[0]

	for _, student := range students {
		if student.Average() > bestAvgGradeStd.Average() {
			bestAvgGradeStd = student
		}
	}

	fmt.Printf("best result %v for %v", bestAvgGradeStd.Average(), bestAvgGradeStd.Name)
}
