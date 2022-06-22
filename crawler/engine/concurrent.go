package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (c ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	c.Scheduler.Run()
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Item: %v", item)
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
