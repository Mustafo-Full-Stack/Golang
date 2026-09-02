package main

import "fmt"

const Pi = 3.14159
const AppName = "MyApp"

/* или лучший вариант(best practice)
const (
	Pi = 3.14159
	AppName = "MyApp"
)
*/

func main() {
	fmt.Println(Pi, AppName)
	// Pi = 3        // ОШИБКА: константу менять нельзя
}
