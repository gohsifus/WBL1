package main

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	value int
	sync.Mutex
}

func (s *safeCounter) inc() {
	s.Lock()
	s.value++
	s.Unlock()
}

func (s safeCounter) String() string {
	return fmt.Sprintf("%v", s.value)
}

func main() {
	wg := &sync.WaitGroup{}
	sc := safeCounter{}
	//1000 го рутин каждая из которых увеличивает счетчик 3 раза
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 3; i++ {
				sc.inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	//Проверить безопасность счетчика можно запустив программу с флагом -race
	//Также о его безопасности говорит ожидаемый результат
	fmt.Println(sc)
}
