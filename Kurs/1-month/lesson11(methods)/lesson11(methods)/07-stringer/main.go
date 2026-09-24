package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Money struct {
	Amount   float64
	Currency string
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", m.Amount, m.Currency)
}

func main() {
	p := Point{X: 3, Y: 4}
	fmt.Println(p)
	fmt.Printf("%v\n", p)

	points := []Point{{1, 2}, {3, 4}}
	fmt.Println(points)

	price := Money{Amount: 1500.5, Currency: "TJS"}
	fmt.Println("Цена:", price)
}

// скажи группе: "вы только что реализовали интерфейс, сами того не зная -
// это fmt.Stringer, завтра разберём, как это работает"
