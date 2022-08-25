package main

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

//New Конструктор
func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//sqrt((xB - xA)^2 + (yB - Ya)^2)
//distance вернет расстояние от p до target
func (p Point) distance(target *Point) float64{
	//Hypot(p, q) = p*p + q*q
	return math.Hypot(p.x - target.x, p.y - target.y)
}

func main() {
	p1 := New(0, 0)
	p2 := New(10, 20)

	fmt.Println(p1.distance(p2))
}
