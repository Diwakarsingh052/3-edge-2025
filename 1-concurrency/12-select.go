package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(3 * time.Second)
		get <- "get done"
	}()

	go func() {
		defer wgWorker.Done()
		time.Sleep(1 * time.Second)
		post <- "post done"
	}()

	go func() {
		defer wgWorker.Done()
		put <- "put done"
		put <- "put 2 done"
	}()

	close(done)
	//	fmt.Println(<-get)
	//	fmt.Println(<-post)
	//	fmt.Println(<-put)

	go func() {
		//for i := 0; i < 3; i++ // this version works, if you know exactly the number of values you would receive
		for {

			select {
			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case msg := <-get:
				fmt.Println(msg)
			case msg := <-post:
				fmt.Println(msg)
			case msg := <-put:
				fmt.Println(msg)

			case <-done:
				fmt.Println("all goroutines finished sending")
				return
			}
		}
	}()
	fmt.Println("end of the main")
	wg.Wait()

}
