package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrDivByZero = errors.New("деление на ноль")
	ErrUnknownOp = errors.New("неизвестная операция")
	ErrBadNumber = errors.New("не является числом")
)

func parseNumber(s string) (float64, error) {
	n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, fmt.Errorf("%q: %w", s, ErrBadNumber)
	}
	return n, nil
}

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("%q: %w", op, ErrUnknownOp)
	}
}

func process(input string) (float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, errors.New("формат: <число> <операция> <число>")
	}

	a, err := parseNumber(parts[0])
	if err != nil {
		return 0, fmt.Errorf("первый операнд: %w", err)
	}

	b, err := parseNumber(parts[2])
	if err != nil {
		return 0, fmt.Errorf("второй операнд: %w", err)
	}

	return calculate(a, parts[1], b)
}

func main() {
	inputs := []string{
		"10 + 5",
		"10 / 0",
		"abc * 2",
		"10 ^ 2",
		"7 - 3",
	}

	for _, in := range inputs {
		res, err := process(in)
		switch {
		case errors.Is(err, ErrDivByZero):
			fmt.Printf("%-10s → нельзя делить на ноль\n", in)
		case errors.Is(err, ErrBadNumber):
			fmt.Printf("%-10s → проверьте числа: %v\n", in, err)
		case errors.Is(err, ErrUnknownOp):
			fmt.Printf("%-10s → доступны: + - * /\n", in)
		case err != nil:
			fmt.Printf("%-10s → ошибка: %v\n", in, err)
		default:
			fmt.Printf("%-10s → %.2f\n", in, res)
		}
	}
}
