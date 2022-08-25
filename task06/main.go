package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

//cancelWithData горутина может считывать значения из переданного канала в инструкции select,
//когда значение будет считано горутина завершается
func cancelWithData(done <-chan int) {
	quit := false
	for !quit {
		select {
		case <-done:
			fmt.Println("Exit")
			quit = true
			break
		default:
			fmt.Println("i am doing something")
		}
	}
}

//cancelWithCloseChannel горутина может отслеживать состояние канала - если канал закрывается - горутина завершается.
func cancelWithCloseChannel(done <-chan struct{}) {
	quit := false
	for !quit {
		select {
		case <-done:
			fmt.Println("Exit")
			quit = true
			break
		default:
			fmt.Println("i am doing something")
		}
	}
}

//cancelWithContext горутину можно завершить с помощью контекста
func cancelWithContext(ctx context.Context) {
	quit := false
	for !quit {
		select {
		case <-ctx.Done():
			fmt.Println("Exit")
			quit = true
			break
		default:
			fmt.Println("i am doing something")
		}
	}
}

func main() {
	//Каждая рутина запускается и завершается через секунду

	ctx, cancel := context.WithCancel(context.Background())
	go cancelWithContext(ctx)
	time.Sleep(time.Second * 1)
	cancel()

	time.Sleep(time.Second * 1)

	dataChan := make(chan int)
	go cancelWithData(dataChan)
	time.Sleep(time.Second * 1)
	dataChan <- 1

	time.Sleep(time.Second * 1)

	doneChan := make(chan struct{})
	go cancelWithCloseChannel(doneChan)
	time.Sleep(time.Second * 1)
	close(doneChan)

	time.Sleep(time.Second * 1)
}
