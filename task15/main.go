package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}
*/

import (
	"fmt"
	"strings"
)

/*
1) Для хранения строки используется переменная уровня пакета, которая живет в течении всей работы программы,
а поскольку срез строки использует тот же массив что и исходная большая строка эта большая строка будет храниться в памяти
2) Проблема с кодировкой: конструкция [:100] создаст подстроку из 100 байт а не символов (кириллица 2 байта)
*/

func createHugeString(n int) string{
	return "JGnkfnkjnОПТАЛОТАЛМОолвоатвлおはようございます"
}

func someFunc() string{
	justString := strings.Builder{}
	v := createHugeString(1<<10)

	count := 0
	//range по строке неявно преобразует utf8
	//Накапливать строку можно с помощью strings.Builder
	//Если использовать + из зв неизменяемости строк каждый раз будет не изменяться старая, а создаваться новая строка
	for _, v := range v{
		justString.WriteRune(v)
		count++
		if count == 17{
			break
		}
	}

	//Если возвращать переменную ее время жизни будет ограничено
	return justString.String()
}

func main(){
	justString := someFunc()
	fmt.Println(justString)
}




