package main

import (
	"fmt"
	"sync"
)

type emailWorker struct {
	name        string
	num_worker  int
	email_queue <-chan int
	wg          *sync.WaitGroup
}

func NewEmailWorker(name string, num_worker int, email_queue <-chan int, wg *sync.WaitGroup) *emailWorker {
	return &emailWorker{
		name:        name,
		num_worker:  num_worker,
		email_queue: email_queue,
		wg:          wg,
	}
}

func (e *emailWorker) RunPool() {
	for i := 0; i < e.num_worker; i++ {
		e.wg.Add(1)
		go func() {
			defer e.wg.Done()
			e.Work()
		}()
	}
}

func (e *emailWorker) Work() {
	for data := range e.email_queue {
		fmt.Println("Receive email with id", data)
	}
}
