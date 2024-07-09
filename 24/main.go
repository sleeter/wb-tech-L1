package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
// Point - струтура с двумя полями x и y
type Point struct {
	x float64
	y float64
}

// newPoint - конструктор
func newPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// distance - метод для рассчета расстояния между точками
func (p Point) distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}
func main() {
	p := newPoint(0.0, 0.0)
	p2 := newPoint(5.0, 5.0)

	fmt.Println(p.distance(p2))
}
