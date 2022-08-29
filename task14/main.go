package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func checkByTypeAssertion(inp interface{}){
	//Проверить тип можно с помощью декларации типов
	if _, ok := inp.(int); ok {
		fmt.Println("i am int")
	}

	//Для удобства можно использовать переключатель типов
	switch inp.(type) {
	case int:
		fmt.Println("I am int!")
	case string:
		fmt.Println("I am string!")
	case bool:
		fmt.Println("I am bool")
	case chan int, chan string, chan interface{}, chan bool:
		fmt.Println("I am chan")
	default:
		fmt.Println("unknown type")
	}
}

func checkByReflect(inp interface{}) {
	inpType := reflect.TypeOf(inp).Kind()

	switch inpType {
	case reflect.Int:
		fmt.Println("i am int!")
	case reflect.String:
		fmt.Println("i am string!")
	case reflect.Bool:
		fmt.Println("i am bool!")
	case reflect.Chan:
		fmt.Println("i am channel!")
	}
}

func main() {
	var whoIAm interface{}
	whoIAm = 23

	checkByTypeAssertion(whoIAm)
	fmt.Println("**********************")
	checkByReflect(whoIAm)
}
