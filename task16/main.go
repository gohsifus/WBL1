package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

import "fmt"

/*
	1. Выбираем любой элемент массив - опорный элемент
	2. Перераспределяем элементы в массиве так чтобы элементы меньшие опорного были слева, а большие справа
	3. После 1 и 2 шага место для опорного элемента найдено
	4. Рекурсивно запускаем шаг 1 и 2  для левой части массива и для правой (опорный элемент - разделитель)
*/

func quickSort(arr []int) []int {
	ret := []int{}

	//База рекурсии, если массив состоит из 1 или меньше - сортировать не надо
	if len(arr) < 2 {
		return arr
	}

	//Опорное значение
	pivot := 0

	left := 1
	right := len(arr) - 1

	for {
		//пока в левой части не найден элемент больше опорного
		if arr[left] <= arr[pivot] {
			left++
		}

		//пока в правой части не найден элемент меньше опорного
		if arr[right] > arr[pivot] {
			right--
		}

		//когда указатели встретились итерация окончена и место для pivot найдено
		if right < left {
			arr[right], arr[pivot] = arr[pivot], arr[right]

			//рекурсивно сортируем левую сторону
			ret = append(ret, quickSort(arr[:right])...)
			//для этого элемента место найдено т.к слева от него гарантированно все меньше а справа все больше
			ret = append(ret, arr[right])
			//рекурсивно сортируем правую сторону
			ret = append(ret, quickSort(arr[right+1:])...)

			return ret
		}

		//когда найдены оба неправильных элемента - меняем
		if arr[left] > arr[pivot] && arr[right] < arr[pivot] {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
}

func main() {
	arr := []int{10, 4, 2, 14, 67, 2, 11, 33, 1, 15}
	fmt.Println(quickSort(arr))
}


//https://prog-cpp.ru/sort-fast/
//https://ru.wikipedia.org/wiki/%D0%91%D1%8B%D1%81%D1%82%D1%80%D0%B0%D1%8F_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B0#%D0%A0%D0%B0%D0%B7%D0%B1%D0%B8%D0%B5%D0%BD%D0%B8%D0%B5_%D0%A5%D0%BE%D0%B0%D1%80%D0%B0