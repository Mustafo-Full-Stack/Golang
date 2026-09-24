package main

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
)

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

type Temperatures []Celsius

func (t Temperatures) Average() Celsius {
	if len(t) == 0 {
		return 0
	}
	var sum Celsius
	for _, v := range t {
		sum += v
	}
	return sum / Celsius(len(t))
}

func main() {
	c := Celsius(25)
	fmt.Println(c, "=", c.ToFahrenheit())

	temps := Temperatures{20, 25, 30}
	fmt.Println("Среднее:", temps.Average())
}
