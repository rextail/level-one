package t18

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна
//выводить итоговое значение счетчика.

//Два варианта: атомики и мьютексы. Мьютекс будет делать гораздо больше, чем нам нужно. Оптимальнее будет взять атомик.

var counter atomic.Uint32

func Incrementer() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			for j := 0; j < a; j++ {
				increment()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.Load())
}

func increment() {
	counter.Add(1)
}
