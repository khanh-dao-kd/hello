package main

import "sync"

/*
product_queue -> product_worker -> email_queue -> email_worker
								-> sms_queue -> sms_worker
*/

func main() {
	var product_wg, email_wg, sms_wg sync.WaitGroup

	product_queue := make(chan int)
	email_queue := make(chan int)
	sms_queue := make(chan int)

	product_worker_pool := NewProductWorker("product_wp", 5, product_queue, []chan<- int{email_queue, sms_queue}, &product_wg)
	email_worker_pool := NewEmailWorker("email_wp", 3, email_queue, &email_wg)
	sms_worker_pool := NewSmsWorker("sms_wp", 3, sms_queue, &sms_wg)

	product_worker_pool.RunPool()
	email_worker_pool.RunPool()
	sms_worker_pool.RunPool()

	for data := 1; data <= 5; data++ {
		product_queue <- data
	}

	close(product_queue)
	product_wg.Wait()

	close(email_queue)
	close(sms_queue)

	email_wg.Wait()
	sms_wg.Wait()
}
