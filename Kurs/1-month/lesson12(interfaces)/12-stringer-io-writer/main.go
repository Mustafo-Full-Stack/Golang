package main

import (
	"fmt"
	"os"
	"strings"
)

// fmt.Stringer - интерфейс из одного метода String() string.
// Если тип его реализует, fmt.Println и fmt.Printf с %v сами вызовут String().
type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Круг радиусом %.1f", c.Radius)
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println(c) // напечатает "Круг радиусом 5.0", а не {5}

	// io.Writer - интерфейс из одного метода Write([]byte) (int, error).
	// os.Stdout, файл, сетевое соединение, буфер в памяти - всё это io.Writer.
	// Одна и та же функция fmt.Fprintln умеет писать в любой из них.
	var sb strings.Builder

	fmt.Fprintln(os.Stdout, "Это ушло в консоль")
	fmt.Fprintln(&sb, "Это ушло в буфер")

	fmt.Println("Из буфера:", sb.String())

	// На неделе про HTTP это пригодится: http.ResponseWriter тоже io.Writer,
	// и в него можно писать точно так же, как в консоль или в буфер.
}
