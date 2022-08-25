package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
*/

import (
	"fmt"
	"strings"
)

//Set Ключи в мапе уникальны
type Set map[string]struct{}

func (s Set) Add(value string){
	s[value] = struct{}{}
}

//Метод для форматирования вывода значений
//Если тип реализует интерфейс Stringer fmt вызывает метод string например при fmt.Println
func (s Set) String() string{
	ret := strings.Builder{}
	ret.WriteString("set[ ")
	for k, _ := range s{
		ret.WriteString(fmt.Sprintf("%s ", k))
	}
	ret.WriteRune(']')
	return ret.String()
}

func main(){
	set := make(Set)
	sequence := []string{"cat", "cat", "dog", "tree"}
	for _, v := range sequence{
		set.Add(v)
	}
	fmt.Println(set)
}
