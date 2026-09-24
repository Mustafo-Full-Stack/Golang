package main

import "fmt"

func main() {
	// var (без make и без литерала) создаёт nil-map
	var broken map[string]int

	fmt.Println(broken == nil, len(broken)) // true 0
	fmt.Println(broken["что-то"])           // 0 — читать из nil-map можно

	// broken["ключ"] = 1
	// раскомментируй строку выше и запусти - будет паника:
	// panic: assignment to entry in nil map
	// это потому что nil-map не готова к записи, в отличие от make(map[...]...)

	// правильно:
	fixed := make(map[string]int)
	fixed["ключ"] = 1
	fmt.Println(fixed)

	// поломка: срез как ключ map — так низя, компилятор не даст
	// m := map[[]string]int{}
	// invalid map key type []string
	// ключом может быть только то, что сравнимо через == (строки, числа, bool, ...)
}
