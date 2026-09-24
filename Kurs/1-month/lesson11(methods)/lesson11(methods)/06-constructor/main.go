package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name     string
	Age      int
	Email    string
	tags     []string
	settings map[string]string
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	if age < 0 || age > 150 {
		return nil, fmt.Errorf("некорректный возраст: %d", age)
	}
	return &User{
		Name: name,
		Age:  age,
		tags: []string{},
		// settings: make(map[string]string), // без этого settings был бы nil-map
	}, nil
}

func (u *User) AddTag(tag string) {
	u.tags = append(u.tags, tag)
}

func (u *User) SetSettings(sts map[string]string) {
	u.settings = sts
}

func main() {
	u, err := NewUser("Али", 100)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	u.AddTag("студент")
	u.AddTag("golang")
	fmt.Printf("%+v\n", *u)

	sts := map[string]string{
		"1": "set1",
		"2": "set2",
	}

	fmt.Printf("%T\n", u.settings)

	u.SetSettings(sts)

	fmt.Println(u)

	_, err = NewUser("", 25)
	fmt.Println("Проверка валидации:", err) // Проверка валидации: имя не может быть пустым
}

// главная польза конструктора - гарантия, что settings не nil-map (вспомни панику из урока 6)
