package main

import (
	"fmt"
	"sync"
)

type smsWorker struct {
	name        string
	num_worker  int
	email_queue <-chan int
	wg          *sync.WaitGroup
}

func NewSmsWorker(name string, num_worker int, email_queue <-chan int, wg *sync.WaitGroup) *smsWorker {
	return &smsWorker{
		name:        name,
		num_worker:  num_worker,
		email_queue: email_queue,
		wg:          wg,
	}
}

func (s *smsWorker) RunPool() {
	for i := 0; i < s.num_worker; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.Work()
		}()
	}
}

func (s *smsWorker) Work() {
	for data := range s.email_queue {
		fmt.Println("Receive sms with id", data)
	}
}
