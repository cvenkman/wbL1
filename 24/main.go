package main

// Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func newPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (a* Point) FindDistance(b *Point) float64 {
	// считаем катеты, берем модуль
	ac := math.Abs(float64(b.x - a.x))
	bc := math.Abs(float64(b.y - a.y))
	// теорема Пифагора
	ab := math.Abs(math.Sqrt(float64(ac * ac + bc * bc)))
	return ab
}

func main() {
	a := newPoint(4, 9)
	b := newPoint(12, 16)

	fmt.Println(a.FindDistance(b)) //10.63014581273465
	fmt.Println(b.FindDistance(a))
}