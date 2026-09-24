package main

import "fmt"

// Урок 10, Шаг 7: срез структур - самый частый паттерн в реальном коде

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	users := []User{
		{Name: "Али", Age: 25, Email: "ali@mail.tj"},
		{Name: "Вера", Age: 30, Email: "vera@mail.tj"},
		{Name: "Тимур", Age: 19, Email: "timur@mail.tj"},
	}

	for i, u := range users {
		fmt.Printf("%d. %-8s %3d лет  %s\n", i+1, u.Name, u.Age, u.Email)
	}
	// 1. Али       25 лет  ali@mail.tj
	// 2. Вера      30 лет  vera@mail.tj
	// 3. Тимур     19 лет  timur@mail.tj

	sum := 0
	for _, u := range users {
		sum += u.Age
	}
	fmt.Printf("Средний возраст: %.1f\n", float64(sum)/float64(len(users))) // 24.7

	// ЛОВУШКА: range отдаёт копию элемента
	for _, u := range users {
		u.Age = 100 // меняем копию, до среза это не доходит
	}
	fmt.Println(users[0].Age) // 25

	// правильно - обращаться к элементу по индексу
	for i := range users {
		users[i].Age = 100
	}
	fmt.Println(users[0].Age) // 100
}

// эту ловушку покажи обязательно: компилятор молчит, программа работает,
// а данные не меняются. одна из самых частых ошибок новичков
