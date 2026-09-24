package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string) error
}

type Sender interface {
	Send(obj any) error
}

type AccountService interface {
	Deposit(amount float64) error
	WithDraw(amount float64) error
	GetBalance() (float64, error)
}

// ReadWriter встраивает Reader и Writer - это не наследование, а сложение
// требований. Чтобы удовлетворять ReadWriter, тип должен уметь и то, и другое.
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	name string
	data string
}

// Оба метода написаны с pointer receiver - иначе Write не сможет менять f.data.
func (f *File) Read() string { return f.data }

func (f *File) Write(s string) error {
	f.data += s
	return nil
}

func main() {
	var rw ReadWriter = &File{name: "test.txt"}
	rw.Write("Привет ")
	rw.Write("мир")
	fmt.Println(rw.Read())

	// Точно так же устроен io.ReadWriter в стандартной библиотеке -
	// он тоже просто объединяет io.Reader и io.Writer.
}
