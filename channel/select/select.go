package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	hasValue := false
	n:=0
	//非阻塞式的接收chan
	for {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
			fmt.Println("from C1:", n)
		case n = <-c2:
			fmt.Println("from C2:", n)
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}
