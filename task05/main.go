package main

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	n := flag.Int("n", 2, "input num of seconds")
	flag.Parse()

	ch := make(chan int)
	//time.After отправит текущее время в канал to после истечения времени
	to := time.After(time.Duration(*n) * time.Second)

	//Отправка значений в канал
	go func() {
		for {
			ch <- rand.Intn(300)
		}
	}()

	//Считываем значения из канала пока не пройдет n секунд
	for {
		select {
		default:
			fmt.Println(<-ch)
		case <-to:
			fmt.Println("timeout")
			os.Exit(0)
		}
	}
}
