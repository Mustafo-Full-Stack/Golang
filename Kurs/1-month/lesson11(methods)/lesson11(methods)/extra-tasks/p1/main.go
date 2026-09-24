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

	for _, grade := range s.Grades {
		total += grade
	}

	return float64(total) / float64(len(s.Grades))
}

func main() {
	students := []Student{
		{
			Name:   "Ali",
			Grades: []int{5, 8, 9, 10, 8},
		},
		{
			Name:   "Bob",
			Grades: []int{2, 6, 1, 10, 9},
		},
		{
			Name:   "John",
			Grades: []int{10, 10, 9, 9, 10},
		},
		{
			Name:   "Kate",
			Grades: []int{8, 8, 8, 9, 9},
		},
		{
			Name:   "Sara",
			Grades: []int{7, 8, 7, 8, 8},
		},
	}

	best := students[0]

	for _, student := range students {
		if student.Average() > best.Average() {
			best = student
		}
	}

	fmt.Printf("Староста: %s\n", best)
	fmt.Printf("Средний балл: %.2f\n", best.Average())
}
