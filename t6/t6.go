package t6

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

//Все способы остановки горутины сводятся к тому, чтобы изнутри подать сигнал на ее завершение
//Из распространенных способ это передача канала, появление значение в котором означает сигнал к остановке,
//а также использование контекста, который в конечном счете тоже подаст сигнал к завершению, но умеет кучу всего.
//С паникой и рековером дело иметь не хочется :)

func Stop() {
	wg := &sync.WaitGroup{}

	stop := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	var stopFlag uint32

	wg.Add(4)
	go g1(wg, stop)
	go g2(wg, ctx)
	go g3(wg, 1*time.Second)
	go g4(wg, &stopFlag)

	time.Sleep(2 * time.Second)

	close(stop) // Останавливаем g1
	cancel()    // Останавливаем g2
	atomic.StoreUint32(&stopFlag, 1)
	// Останавливаем g3
	wg.Wait()
}

func g1(wg *sync.WaitGroup, stop <-chan struct{}) {
	//Можно создать канал хоть какого типа,
	//но struct{} хотя бы память не занимает
	defer func() {
		fmt.Println("stopped g1")
		wg.Done()
	}()
	for {
		select {
		case <-stop:
			return
		default:
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func g2(wg *sync.WaitGroup, ctx context.Context) {
	//базовая остановка по отмене контекста
	defer func() {
		fmt.Println("stopped g2")
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func g3(wg *sync.WaitGroup, dur time.Duration) {
	//Зададим TTL для горутины
	defer func() {
		fmt.Println("stopped g3")
		wg.Done()
	}()
	timer := time.After(dur)
	for {
		select {
		case <-timer:
			return
		default:
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func g4(wg *sync.WaitGroup, flag *uint32) {
	//Создадим глобальную флаг-переменную, и по ее значению будем останавливать горутину
	defer func() {
		fmt.Println("stopped g4")
		wg.Done()
	}()
	for {
		if atomic.LoadUint32(flag) == 1 {
			return
		}
	}
}
