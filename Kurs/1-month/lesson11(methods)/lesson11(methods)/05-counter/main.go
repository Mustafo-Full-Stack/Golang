package main

import "fmt"

type Counter struct {
	count int
	step  int
}

func (c *Counter) Inc() {
	c.count += c.step
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c Counter) Value() int {
	return c.count
}

func main() {
	c := Counter{step: 5}
	c.Inc()
	c.Inc()
	c.Inc()
	fmt.Println(c.Value()) // 15

	c.Reset()
	fmt.Println(c.Value()) // 0
}

// поля count и step - с маленькой буквы, менять их можно только через методы.
// это инкапсуляция без ключевого слова private
