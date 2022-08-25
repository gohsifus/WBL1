package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func strReverse(inp string) string {
	//Преобразуем строку в массив rune чтобы работать с символами а не с байтами
	chars := []rune(inp)
	//Складывать строки через += нельзя, поэтому используем builder
	builder := strings.Builder{}

	for i := len(chars) - 1; i >= 0; i-- {
		builder.WriteRune(chars[i])
	}

	return builder.String()
}

func strReverse2(inp string) string {
	i := utf8.RuneCountInString(inp) - 1
	ret := make([]rune, i + 1)

	//range по строке возвращает руны а не байты
	for _, v := range inp {
		ret[i] = v
		i--
	}

	return string(ret)
}

func main() {
	fmt.Println(strReverse("qweПривет"))
	fmt.Println(strReverse2("qweПривет"))
}
