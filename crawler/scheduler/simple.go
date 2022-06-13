package scheduler

import "gonb/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s SimpleScheduler) Submit(request engine.Request) {
	//send request down to worker chan
	s.WorkerChan <- request
}

