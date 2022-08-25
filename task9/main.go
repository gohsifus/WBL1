package main

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x^2,
после чего данные из второго канала должны выводиться в stdout.
*/

import "fmt"

//С помощью каналов go рутины можно организовать так чтобы выход одной был входом для другой - это и называется конвейером

func main() {
	data := [12]int{1, 2, 3, 4, 6, 5, 7, 8, 67, 9, 8, 9}
	stage1 := make(chan int)
	stage2 := make(chan int)

	go func() {
		for _, v := range data {
			stage1 <- v
		}
		//С помощью закрытия канала можно сообщить следующей стадии что данные переданы
		close(stage1)
	}()

	go func() {
		for v := range stage1{
			stage2 <- v * v
		}
		close(stage2)
	}()

	for {
		val, ok := <-stage2
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
