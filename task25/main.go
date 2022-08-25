package main

/*
Реализовать собственную функцию sleep
*/

import (
	"fmt"
	"time"
)

//sleep time.After возвращает канал, ждет указанное время и по его прошествии записывает текущее время в канао,
//рутина блокируется пока не считает информацию из канала
func sleep(duration time.Duration){
	/*
	Можно использовать таймер
	timer := time.NewTimer(duration)
	<-timer.C
	*/

	<-time.After(duration)
}

//sleep Способ 2 - устанавливаем время после которого надо возобновить работу и цикле проверяем не наступило ли это время
func sleep2(duration time.Duration){
	timeEnd := time.Now().Add(duration)
	for {
		if time.Now().After(timeEnd) {
			break
		}
	}
}

func main(){
	fmt.Println("start")
	sleep(time.Second * 10)
	fmt.Println("end")
}
