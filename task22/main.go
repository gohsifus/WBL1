package main

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
*/

import (
	"fmt"
	"math/big"
)

//Для работы с большими числами есть библиотека big
//int64 хранит до 24e18 при попытке записать число больше будет переполнение
//big позволяет обойти это ограничение
//Кроме big.Int есть big.Float и big.Rat рациональные числа(которые можно представить в виде отношения двух целых)

func div(a, b *big.Int) *big.Int{
	ret := big.Int{}
	return ret.Div(a, b)
}

func add(a, b *big.Int) *big.Int{
	ret := big.Int{}
	return ret.Add(a, b)
}

func mul(a, b *big.Int) *big.Int{
	ret := big.Int{}
	return ret.Mul(a, b)
}

func sub(a, b *big.Int) *big.Int{
	ret := big.Int{}
	return ret.Sub(a, b)
}

func main(){
	a := &big.Int{}
	b := &big.Int{}

	//Для инициализации можно использовать строки
	a.SetString("2400000000000000000000000000000", 10)
	b.SetString("2400000000000000", 10)

	fmt.Println(div(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(add(a, b))
}
