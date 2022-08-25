package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/

import (
	"fmt"
	"unicode"
)

func isUnique(inp string) bool {
	m := make(map[rune]struct{})

	for _, v := range inp {
		//Записываем в мапу руны в нижнем регистре
		m[unicode.ToLower(v)] = struct{}{}
	}

	//Если кол-во не совпадает значит какой-то символ записался в мапу 2 раза и он не уникален
	return len(m) == len(inp)
}

func main() {

	str := "qwe12"
	fmt.Println(isUnique(str))
}
