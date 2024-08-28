package t2

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

//Если мы просто на каждое число будем создавать по горутине, в таком случае данн

//Мои мысли: если реализовывать простенький пайплайн, в котором в одной горутине
//отправляем в канал, а во второй читаем из него и пишем в stdout, то горутина, ждущая из канала данные,
//будет все время внаглую засыпать, потом какое-то время раскачиваться и только затем работать.
//Если хотим написать такой код, который мог бы выполняться параллельно, тогда придется
//как-то гарантировать, что пишущая в stdout горутина не будет туда-сюда парковаться, иначе толку не будет.

func PrintSquaresV1(arr []int) {
	if len(arr) == 0 {
		return
	}

	var out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	var wg sync.WaitGroup

	numCh := make(chan int)
	go func() {
		for _, v := range arr {
			numCh <- v * v
		}
	}()
	wg.Add(len(arr))
	go func() {
		for num := range numCh {
			fmt.Fprintln(out, num)
			wg.Done()
		}
	}()
	wg.Wait()
	close(numCh)
}

//Можем последовать примеру Sameer Ajumani'a, опубликованному в Go Blog в 2014

func PrintSquaresV2(arr []int) {
	var out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	numsCh := seqToChan(arr...)
	squareCh := square(numsCh)

	for n := range squareCh {
		fmt.Fprintln(out, n)
	}

}

//Хорошо, а если мы теперь возьмем, и заставим две горутины читать из канала с числами:
//пусть наберут себе данных сколько успеют, возведут их себе спокойно в квадрат,
//а мы потом объединим результаты их усилий в один канал и прочитаем из него? (fan-in/fan-out)

func PrintSquaresV3(arr []int) {
	var out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	numsCh := seqToChan(arr...)

	//Порядок не гарантирован

	squareCh1 := square(numsCh)
	squareCh2 := square(numsCh)

	for n := range mergeChans(squareCh1, squareCh2) {
		fmt.Fprintln(out, n)
	}
}

func seqToChan(nums ...int) chan int {
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
	//кол-во каналов, которые пишут в канал
	wg.Add(len(chs))
	go func() {
		for _, ch := range chs {
			go output(ch)
		}
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
