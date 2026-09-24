package main

import "fmt"

// Урок 10, Шаг 10: структуры внутри map - и главная ловушка

type User struct {
	Name string
	Age  int
}

func main() {
	catalog := map[string]User{
		"admin": {Name: "Админ", Age: 40},
	}
	fmt.Println(catalog["admin"].Name) // Админ - читать можно спокойно

	// ЛОВУШКА: менять поле напрямую нельзя
	// catalog["admin"].Age = 41
	// cannot assign to struct field catalog["admin"].Age in map

	// решение 1 - достать, поменять, положить обратно
	u := catalog["admin"]
	u.Age = 41
	catalog["admin"] = u
	fmt.Println(catalog["admin"].Age) // 41

	// решение 2 (обычно лучше) - хранить указатели
	catalog2 := map[string]*User{
		"admin": {Name: "Админ", Age: 40},
	}
	catalog2["admin"].Age = 41
	fmt.Println(catalog2["admin"].Age) // 41
}

// раскомментируй строку с ошибкой вживую - такое лучше один раз увидеть.
// причина: map отдаёт копию значения, и у этой копии нет адреса,
// значит присвоить в её поле некуда
