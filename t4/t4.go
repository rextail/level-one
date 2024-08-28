package t4

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
//произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//Выбрать и обосновать способ завершения работы всех воркеров.

func WorkerPool(nWorkers int) {
	wg := &sync.WaitGroup{}

	stop := make(chan struct{})

	numbers := gen(stop)

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go isEven(wg, numbers)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(stop)
	wg.Wait()
}

func gen(stop chan struct{}) chan int {
	//генерируем какие-то данные, чтобы воркерам было чем заняться
	out := make(chan int)
	go func() {
		defer func() {
			fmt.Println("numbers chan was closed")
			close(out)
		}()
		for {
			select {
			case <-stop:
				//У нас N воркеров обрабатывают данные из общего канала.
				//Закрытие общего канала прекратит поток данных воркерам, когда канал опустеет, воркеры завершат работу
				return
			default:
				out <- rand.Int()
			}
		}
	}()
	return out
}

func isEven(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for i := range in {
		fmt.Println(i%2 == 0)
	}
}
