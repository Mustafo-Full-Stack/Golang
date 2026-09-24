package main

import "fmt"

// Урок 10, Шаг 6: указатель на структуру - как менять оригинал

type User struct {
	Name string
	Age  int
}

func birthday(u User) {
	u.Age++ // меняет копию
}

func birthdayPtr(u *User) {
	u.Age++ // меняет оригинал
}

func main() {
	user := User{Name: "Али", Age: 25}

	birthday(user)
	fmt.Println("После birthday:", user.Age) // 25

	birthdayPtr(&user)
	fmt.Println("После birthdayPtr:", user.Age) // 26

	// указатель на структуру - очень частый приём
	// p := User{Name: "Вера", Age: 30}
	p := &User{Name: "Вера", Age: 30}
	fmt.Printf("%T\n", p)
	fmt.Println(p.Name) // Вера

	p.Age = 31

	fmt.Printf("%+v\n", *p) // {Name:Вера Age:31}
}

// подчеркни: писать (*p).Name не нужно - Go разыменовывает указатель сам.
// это просто синтаксическое удобство, никакой магии
