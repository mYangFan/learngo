package main

import (
	"fmt"
	"reflect"
	"time"
)

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func worker(id int, c chan int) {
	for  n := range c{
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)

	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//channelClose()
	a := "abc"

	field := reflect.TypeOf(a)
	value := reflect.ValueOf(a)
	fmt.Println(field)
	fmt.Println(value)
}
