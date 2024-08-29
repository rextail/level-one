package t25

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep

func Sleep1(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	fmt.Printf("sleep1: woke up after %v seconds \n", sec)
}

func Sleep2(sec int) {
	<-time.After(time.Second * time.Duration(sec))
	fmt.Printf("sleep2: woke up after %v seconds \n", sec)
}
