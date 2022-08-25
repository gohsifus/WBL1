package main
/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

import "fmt"

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров
*/

import (
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// worker в бесконечном цикле считывает данные с канала
// если канал закрыт worker прекращает работу
func worker(input <-chan int) {
	for {
		if val, ok := <-input; !ok {
			fmt.Println("Завершение чтения")
			break
		} else {
			fmt.Println(val)
		}
	}
}

func main() {
	//"Необходима возможность выбора количества воркеров при старте"
	n := flag.Int("n", 5, "input num of workers")
	flag.Parse()

	streamChan := make(chan int)
	quit := false

	//Отслеживание Ctr+c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	//Запускаем n воркеров
	for i := 0; i < *n; i++ {
		go worker(streamChan)
	}

	//"Реализовать постоянную запись данных в канал (главный поток)"
	for !quit {
		select {
		//Если Ctrl + c
		case <-ch:
			fmt.Println("Завершение записи")
			close(streamChan)
			fmt.Println("Завершение программы")
			quit = true
		default:
			streamChan <- rand.Intn(300)
		}
	}
}
