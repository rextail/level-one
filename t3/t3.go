package t3

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.

//Так как последовательность квадратов чисел при суммировании не играет никакой роли, можем
//все это дело через fan-in/fan-out спокойно посчитать.

func PrintSquaresSum(nums ...int) {

	numsCh := seqToChan(nums)

	//Порядок не гарантирован

	squareCh1 := square(numsCh)
	squareCh2 := square(numsCh)

	ans := 0
	for n := range mergeChans(squareCh1, squareCh2) {
		ans += n
	}
	fmt.Println(ans)
}

func seqToChan(nums []int) chan int {
	out := make(chan int)
	//канал будет заполняться в другой горутине
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	//возвращаем канал, в который будут поступать переданные данные
	return out
}

func square(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func mergeChans(chs ...<-chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	//функция для обработки значений из канала
	output := func(c <-chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(chs))
	go func() {
		for _, ch := range chs {
			go output(ch)
		}
	}()
	go func() {
		//ждем, пока все каналы, из которых мы читаем, закроются
		wg.Wait()
		close(out)
	}()
	return out
}
