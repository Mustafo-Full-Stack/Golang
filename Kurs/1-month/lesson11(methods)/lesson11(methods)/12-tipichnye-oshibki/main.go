package main

import "fmt"

type Counter struct {
	count int
}

// IncBroken - специально с value-receiver, только для сравнения с Inc ниже.
// в реальном коде так смешивать receiver'ы одного типа нельзя
func (c Counter) IncBroken() {
	c.count++
}

func (c *Counter) Inc() {
	c.count++
}

func main() {
	c := Counter{}
	c.IncBroken()
	c.IncBroken()
	fmt.Println(c.count) // 0 - компилируется, но не работает. самая опасная ошибка - без единого сообщения

	c.Inc()
	fmt.Println(c.count) // 1 - так правильно

	// поломка 2 - метод на чужом (не своём) типе
	// func (i int) Double() int { return i * 2 }
	// cannot define new methods on non-local type int
	// решение - свой тип: type MyInt int

	// поломка 3 - pointer-метод у элемента map
	m := map[string]Counter{"a": {}}
	// m["a"].Inc()
	// cannot call pointer method Inc on m["a"]
	fmt.Println(m["a"].count) // 0 - читать можно, менять и вызывать pointer-метод - нет. решение: map[string]*Counter
}

// поломка 4 - бесконечная рекурсия в String()
// func (u User) String() string {
// 	return fmt.Sprintf("%v", u) // %v внутри String() снова вызывает String() - stack overflow
// }
// правильно - обращаться к полям напрямую: fmt.Sprintf("%s, %d", u.Name, u.Age)

// поломка 1 - самая коварная: компилятор молчит.
// скажи группе: "если данные не меняются, а код выглядит правильно -
// первым делом проверьте receiver"
