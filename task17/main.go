package main

/*
Реализовать бинарный поиск встроенными методами языка.
*/

import (
	"fmt"
	"sort"
)

/*
	Бинарный поиск происходит в отсортированном массиве

	1. Искомое значение сравнивается с средним в массиве - если равны то поиск окончен
	2. Если не равны то шаг 1 повторяется для левой или правой части в зависимости от результата сравнения на шаге 1
*/
func binSearch(slice []int, target int, left, right int) (int, bool) {
	//left и right - первый и последний индексы части массива в которой ведется поиск
	middle := (right-left)/2 + left
	//fmt.Println(fmt.Sprintf("middle: %v, left: %v, right: %v", middle, left, right))

	//Если указатели встретились значит остался 1 элемент
	if right == left {
		return right, target == slice[right]
	}

	if target == slice[middle] {
		//Элемент найден
		return middle, true
	} else if target < slice[middle] {
		//Ищем в левой части
		right = middle - 1
		return binSearch(slice, target, left, right)
	} else {
		//Ищем в правой части
		left = middle + 1
		return binSearch(slice, target, left, right)
	}
}

func binSearch2(slice []int, target int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		middle := (right - left) / 2 + left
		//fmt.Println(fmt.Sprintf("middle: %v, left: %v, right: %v", middle, left, right))

		if target == slice[middle] {
			return middle
		} else if target < slice[middle]{
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	//1 2 12 23 34 354 455 6565
	slice := []int{1, 12, 34, 2, 354, 6565, 23, 455, 12}
	sort.Ints(slice)
	fmt.Println(slice)

	if index, ok := binSearch(slice, 455, 0, len(slice)-1); ok {
		fmt.Printf("Индекс искомого элемента: %v\n", index)
	}

	fmt.Printf("Индекс искомого элемента 2 реализация: %v", binSearch2(slice, 455))
}
