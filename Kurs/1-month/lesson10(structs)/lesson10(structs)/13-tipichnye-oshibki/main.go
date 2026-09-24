package main

import "fmt"

// Урок 10: "сломай специально" - четыре типичные ошибки со структурами

type User struct {
	Name string
	Age  int
}

type Group struct {
	Members []string
}

func main() {
	users := []User{
		{Name: "Али", Age: 25},
		{Name: "Вера", Age: 30},
	}

	// поломка 1 - изменение структуры, лежащей в map
	m := map[string]User{"a": {Name: "Али"}}
	// m["a"].Age = 30
	// cannot assign to struct field m["a"].Age in map
	fmt.Println(m["a"].Name) // Али

	// поломка 2 - изменение в range.
	// компилируется, запускается, ничего не меняет - и ни одного сообщения
	for _, u := range users {
		u.Age++
	}
	fmt.Println(users[0].Age) // 25, а не 26

	// поломка 3 - инициализация по порядку после правки структуры.
	// добавь в User поле Phone string между Name и Age и раскомментируй строку:
	// bad := User{"Али", 25}
	// cannot use 25 (untyped int constant) as string value in struct literal
	// too few values in struct literal of type User
	// мораль: всегда указывай имена полей - тогда правка структуры ничего не ломает

	// поломка 4 - сравнение структур, внутри которых срез
	g1 := Group{}
	g2 := Group{}
	// fmt.Println(g1 == g2)
	// invalid operation: g1 == g2 (struct containing []string cannot be compared)
	fmt.Println(len(g1.Members), len(g2.Members)) // 0 0
}

// поломка 2 - самая коварная из всех: компилятор молчит.
// скажи группе: "если данные не меняются, а код выглядит правильным -
// первым делом проверьте, не меняете ли вы копию из range"
