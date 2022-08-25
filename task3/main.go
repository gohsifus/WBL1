package main

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

//way1 организуем потокобезопасное обращение к переменной с помощью мьютекса
func way1(data []int) int {
	wg := &sync.WaitGroup{}
	var (
		mu    sync.Mutex
		value int
	)

	for _, v := range data {
		wg.Add(1)
		go func(val int) {
			square := val * val
			mu.Lock()
			value += square
			mu.Unlock()
			wg.Done()
		}(v)
	}

	wg.Wait()
	return value
}

//way2 так как chan потокобезопасный можем использовать его
//каждая рутина после вычислений записывает значение в chan
//после запуска рутин считываем значения n раз и формируем сумму
func way2(data []int) int {
	sum := 0
	ch := make(chan int)

	for _, v := range data {
		go func(val int) {
			square := val * val
			ch <- square
		}(v)
	}

	for range data {
		sum += <-ch
	}

	return sum
}

func main() {
	sequence := []int{2, 4, 6, 8, 10}
	fmt.Println(way1(sequence))
	fmt.Println("***********")
	fmt.Println(way2(sequence))
}
