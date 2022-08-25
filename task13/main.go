package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main(){
	i := 1
	j := 2
	fmt.Println(i, j)
	//Прежде чем любая из переменных в левой части получит новое значение, вычисляются все выражения из правой части
	//поэтому можно легко поменять местами
	j, i = i, j
	fmt.Println(i, j)
}