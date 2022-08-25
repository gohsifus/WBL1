package main

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

import (
	"fmt"
	"strings"
)

func revertStringByWords(inp string) string{
	words := strings.Split(inp, " ")
	builder := strings.Builder{}

	i := len(words) - 1
	for range words{
		builder.WriteString(words[i])
		builder.WriteRune(' ')
		i--
	}

	return builder.String()
}

//Второй способ без выделения памяти для builder
func revertStringByWords2(inp string) string{
	words := strings.Split(inp, " ")

	right := len(words) - 1
	for left := 0; left < len(words) / 2; left++{
		words[left], words[right] = words[right], words[left]
		right--
	}

	return strings.Join(words, " ")
}

func main(){
	str := "snow dog sun"

	fmt.Println(revertStringByWords(str))
	fmt.Println(revertStringByWords2(str))
}
