package t5

import (
	"bufio"
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func Pipe() {
	out := bufio.NewWriter(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		out.Flush()
	}()

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				num := rand.Int()
				ch <- num
				time.Sleep(1 * time.Second)
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case n, ok := <-ch:
			if !ok {
				return
			}
			fmt.Fprintln(out, n)
		}
	}
}
