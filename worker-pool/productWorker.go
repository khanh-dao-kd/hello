package main

import "sync"

type productWorker struct {
	name              string
	num_worker        int
	product_queue     <-chan int
	list_of_des_queue []chan<- int
	wg                *sync.WaitGroup
}

func NewProductWorker(name string, num_worker int, product_queue <-chan int, list_of_des_queue []chan<- int, wg *sync.WaitGroup) *productWorker {
	return &productWorker{
		name:              name,
		num_worker:        num_worker,
		product_queue:     product_queue,
		list_of_des_queue: list_of_des_queue,
		wg:                wg,
	}
}

func (p *productWorker) RunPool() {
	for i := 0; i < p.num_worker; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.Work()
		}()
	}
}

func (p *productWorker) Work() {
	for data := range p.product_queue {
		for _, des_queue := range p.list_of_des_queue {
			des_queue <- data
		}
	}
}
