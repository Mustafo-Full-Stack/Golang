package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("=== 1. Порядок выполнения defer (LIFO) ===")
	//deferOrder()
	//
	//fmt.Println("\n=== 2. defer f(x) - аргументы вычисляются СРАЗУ, в момент defer ===")
	//deferImmediateArgs()

	fmt.Println("\n=== 3. defer func() {...}() - замыкание читает переменную В МОМЕНТ ВЫПОЛНЕНИЯ ===")
	deferClosure()

	fmt.Println("\n=== 4. Реальный пример: замер времени выполнения функции ===")
	slowOperation()

	//fmt.Println("\n=== 5. Реальный пример: defer + recover через именованный результат ===")
	//result, err := safeDivide(10, 0)
	//fmt.Println("Результат:", result, "| Ошибка:", err)
}

// 1. Несколько defer выполняются в обратном порядке - как стопка тарелок:
// последний положил - первый снял.
func deferOrder() {
	fmt.Println("начало функции")
	defer fmt.Println("defer 1 (объявлен первым)")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3 (объявлен последним)")
	fmt.Println("конец функции")
}

// 2. У defer fmt.Println("x =", x) аргумент x вычисляется прямо здесь,
// в момент объявления defer, и "замораживается". Дальнейшие изменения x
// на уже отложенный вызов не влияют.

func deferImmediateArgs() {
	x := 1
	defer fmt.Println("x в момент defer:", x) // x = 1 зафиксирован уже сейчас
	x = 2
	fmt.Println("x изменили на:", x)
}

// 3. У defer func() {...}() тела функции нет аргументов - она замыкается
// на переменную x и читает её значение только в момент реального выполнения,
// то есть перед выходом из функции.
func deferClosure() {
	x := 1
	defer func() {
		fmt.Println("x в момент выполнения defer:", x) // видит актуальное значение
	}()
	x = 2
	fmt.Println("x изменили на:", x)
}

// Классический паттерн из реальных бэкендов: замер времени работы функции
// (запросы к БД, обработчики HTTP, тяжёлые вычисления). start передаётся
// аргументом - как в пункте 2 он "замораживается" в момент defer, а тело
// функции выполняется позже, как в пункте 3.
func slowOperation() {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("операция заняла:", time.Since(start))
	}(start)

	time.Sleep(50 * time.Millisecond)
}

// Ещё один реальный паттерн: defer с recover, который "спасает" функцию
// от паники и записывает ошибку в именованный результат. Так пишут,
// например, обработчики, которым нельзя ронять всю программу целиком.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("восстановились после паники: %v", r)
		}
	}()
	result = a / b // при b == 0 здесь произойдёт паника
	return result, nil
}
