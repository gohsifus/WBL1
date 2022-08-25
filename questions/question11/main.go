package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func(){
		time.Sleep(time.Second*10)
	}()
	go func(){
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(wg sync.WaitGroup, i int) {
				fmt.Println(i)
				wg.Done()
			}(wg, i)
		}
		wg.Wait()
		fmt.Println("exit")
	}()
	wg2.Wait()
	fmt.Println("qwe")
}
