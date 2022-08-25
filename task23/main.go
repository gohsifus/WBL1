package main

/*
Удалить i-ый элемент из слайса
*/

import "fmt"

func removeElementByIndex(inp []int, i int) []int {
	//append добавляет элемент в конец переданного слайса
	//возвращает полученный слайс
	//если при append cap переполняется, то append выделяет новый массив для слайса,
	//и тогда старый массив не будет меняться
	//например если вместо inp[i+1:] передать inp[:] то cap переполнится и массив лежащий в основе inp не изменится
	return append(inp[:i], inp[i+1:]...)
}
//второй способ с помощью цикла - перезаписываем элементы начиная с index, возвращаем новый слайс без последнего элемента
func removeElementByIndex2(inp []int, index int) []int {
	for i := 0; i < len(inp) - 1; i++{
		if i == index{
			inp[i] = inp[i+1]
			index++
		}
	}
	return inp[:index]
}

func main() {
	slice := []int{1, 2, 3, 3, 54, 5, 3, 53, 4}
	slice = removeElementByIndex(slice, 2)
	slice = removeElementByIndex2(slice, 2)
	fmt.Println(slice)
}
