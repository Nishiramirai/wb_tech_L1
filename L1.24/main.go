package main

import (
	"fmt"
	"math"
)

// Расстояние между точками

// Разработать программу нахождения расстояния между двумя точками на плоскости.
//  Точки представлены в виде структуры Point с инкапсулированными (приватными)
//  полями x, y (типа float64) и конструктором.
//  Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point
//  и метод Distance(other Point) float64.

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(-1, 7)
	p2 := NewPoint(7, 1)
	fmt.Println(p1.Distance(p2))
}
