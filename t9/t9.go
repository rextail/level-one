package t9

import (
	"bufio"
	"fmt"
	"os"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
//результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

//Штош:)

func AnotherPipeline(arr []int) {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	nums := fillByNums(arr)
	squares := squareNums(nums)
	for n := range squares {
		fmt.Fprintln(out, n)
	}
}

func fillByNums(arr []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range arr {
			out <- v
		}
		close(out)
	}()
	return out
}

func squareNums(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}
