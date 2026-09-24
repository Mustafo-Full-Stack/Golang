package main

import "fmt"

type Engine struct {
	Power int
	Type  string
}

func (e Engine) Start() string {
	return fmt.Sprintf("Двигатель %s (%d л.с.) запущен", e.Type, e.Power)
}

func (e Engine) Stop() string {
	return "Двигатель остановлен"
}

type Car struct {
	Engine
	Brand string
	Model string
}

type SportsCar struct {
	Car
	TopSpeed int
}

func (s SportsCar) Start() string {
	return "РЁВ! " + s.Car.Start() // вызываем "родительский" метод и дополняем его
}

func main() {
	s := SportsCar{
		Car:      Car{Engine: Engine{Power: 400, Type: "V8"}, Brand: "Ferrari"},
		TopSpeed: 320,
	}
	fmt.Println(s.Start())
	fmt.Println(s.Stop())
}

// в других языках это назвали бы переопределением - но в Go нет ни ключевых слов, ни иерархий
