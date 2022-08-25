package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

import "fmt"

type Human struct {
	name string
	Age  int
}

func (h Human) speak() string {
	return fmt.Sprintf("Hello, my name is %s, i am %d year old.", h.name, h.Age)
}

// Action
//Структура Human встроена в Action, Action получает доступ к полям и методам Human
//Такая запись позволяет обращаться к методам и полям встроенной структуры без использования промежуточного имени поля,
//Action.Human.name - на самом деле анонимное поле имеет имя Human но его можно опустить при обращении
type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{
			"Nikita",
			28,
		},
	}

	//Экспортируемость полей в главной структуре не влияет на уровень доступа
	fmt.Println(action.name)
	fmt.Println(action.Age)

	fmt.Println(action.speak())
}
