package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	get := make(chan string, 2)
	post := make(chan string, 2)
	put := make(chan string, 2)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		get <- "get done"
	}()

	go func() {
		defer wgWorker.Done()
		post <- "post done"
	}()

	go func() {
		defer wgWorker.Done()
		put <- "put done"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done)
	}()
	//	fmt.Println(<-get)
	//	fmt.Println(<-post)
	//	fmt.Println(<-put)

	wg.Add(1)
	go func() {
		defer wg.Done()
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
