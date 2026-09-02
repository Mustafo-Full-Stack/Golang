package main

import "fmt"

func main() {

	// Объявляем переменные без значений
	var (
		intValue   int
		int8Value  int8
		int16Value int16
		int32Value int32
		int64Value int64

		uintValue   uint
		uint8Value  uint8
		uint16Value uint16
		uint32Value uint32
		uint64Value uint64

		float32Value float32
		float64Value float64

		boolValue bool

		runeValue rune
		byteValue byte

		stringValue string

		complex64Value  complex64
		complex128Value complex128
	)

	fmt.Println("=== Значения по умолчанию ===")

	fmt.Println("int:", intValue)
	fmt.Println("int8:", int8Value)
	fmt.Println("int16:", int16Value)
	fmt.Println("int32:", int32Value)
	fmt.Println("int64:", int64Value)

	fmt.Println("uint:", uintValue)
	fmt.Println("uint8:", uint8Value)
	fmt.Println("uint16:", uint16Value)
	fmt.Println("uint32:", uint32Value)
	fmt.Println("uint64:", uint64Value)

	fmt.Println("float32:", float32Value)
	fmt.Println("float64:", float64Value)

	fmt.Println("bool:", boolValue)

	fmt.Println("rune:", runeValue)
	fmt.Println("byte:", byteValue)

	fmt.Println("string:", stringValue)

	fmt.Println("complex64:", complex64Value)
	fmt.Println("complex128:", complex128Value)

	// Присваиваем значения

	intValue = 100
	int8Value = 8
	int16Value = 1600
	int32Value = 32000
	int64Value = 640000

	uintValue = 500
	uint8Value = 255
	uint16Value = 65000
	uint32Value = 320000
	uint64Value = 6400000

	float32Value = 3.14
	float64Value = 6.28

	boolValue = true

	runeValue = 'G'
	byteValue = 65

	stringValue = "Golang"

	complex64Value = 2 + 3i
	complex128Value = 5 + 7i

	fmt.Println("\n=== После присваивания значений ===")

	fmt.Println("int:", intValue)
	fmt.Println("int8:", int8Value)
	fmt.Println("int16:", int16Value)
	fmt.Println("int32:", int32Value)
	fmt.Println("int64:", int64Value)

	fmt.Println("uint:", uintValue)
	fmt.Println("uint8:", uint8Value)
	fmt.Println("uint16:", uint16Value)
	fmt.Println("uint32:", uint32Value)
	fmt.Println("uint64:", uint64Value)

	fmt.Println("float32:", float32Value)
	fmt.Println("float64:", float64Value)

	fmt.Println("bool:", boolValue)

	fmt.Println("rune:", runeValue)
	fmt.Println("byte:", byteValue)

	fmt.Println("string:", stringValue)

	fmt.Println("complex64:", complex64Value)
	fmt.Println("complex128:", complex128Value)

}
