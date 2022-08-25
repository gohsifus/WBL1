package main

/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
*/

import "fmt"

type Groups map[int][]float32

func (g Groups) Add(value float32) {
	//При делении int дробное значение отбрасывается
	key := (int(value) / 10) * 10
	g[key] = append(g[key], value)
}

func main() {
	g := make(Groups)
	sequence := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 123.4}
	for _, v := range sequence {
		g.Add(v)
	}

	fmt.Println(g)
}
