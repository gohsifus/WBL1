package main

/*
Реализовать конкурентную запись данных в map
*/

import (
	"fmt"
	"math/rand"
	"sync"
)

//Параллельную запись в map можно организовать с помощью mutex
type safeMap struct {
	m map[int]int
	sync.Mutex
}

//Метод для безопасной записи
func (s *safeMap) add(key, value int) {
	s.Lock()
	s.m[key] = value
	s.Unlock()
}

//Для проверки использовался детектор гонки
func main() {
	wg := &sync.WaitGroup{}

	safe := safeMap{
		map[int]int{},
		sync.Mutex{},
	}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			safe.add(rand.Intn(30), rand.Intn(30))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(safe.m)
}
