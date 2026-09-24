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
	Engine // встраивание - Car получает поля и методы Engine бесплатно
	Brand  string
	Model  string
}

func main() {
	c := Car{
		Engine: Engine{Power: 150, Type: "бензиновый"},
		Brand:  "Toyota",
		Model:  "Camry",
	}

	fmt.Println(c.Start())        // Двигатель бензиновый (150 л.с.) запущен - метод Engine доступен напрямую
	fmt.Println(c.Power)          // 150 - поле тоже "поднялось" наверх
	fmt.Println(c.Brand)          // Toyota
	fmt.Println(c.Engine.Start()) // тот же результат - полная форма доступа тоже работает
}

// скажи группе: "Car получил умения Engine бесплатно, просто встроив его -
// это композиция вместо наследования"
