package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main(){
	var whoIAm interface{}
	whoIAm = 23

	//Проверить тип можно с помощью декларации типов
	if x, ok := whoIAm.(int); ok{
		fmt.Println(reflect.TypeOf(x))
	}

	//Для удобства можно использовать переключатель типов
	switch whoIAm.(type){
	case int:
		fmt.Println("I am int!")
	case string:
		fmt.Println("I am string!")
	case bool:
		fmt.Println("I am bool")
	case chan int, chan string, chan interface{}, chan bool:
		fmt.Println("I am chan")
	default:
		//Также тип можно узнать с помощью рефлексии
		fmt.Println(reflect.TypeOf(whoIAm))
	}
}
