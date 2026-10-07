package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"comment,omitempty"`
}

func main() {
	student_1 := Student{
		Name:    "Аня",
		Grade:   25,
		Comment: "",
	}

	student_2 := Student{
		Name:    "Иван",
		Grade:   50,
		Comment: "Топ 1",
	}

	data, err := json.MarshalIndent([]Student{student_1, student_2}, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
