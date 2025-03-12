package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create one more waitgroup, use that to track fanned out goroutines
	// once fanned out goroutines finishes close the channel
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)

		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)
			go func(id int) {
				defer wgWorker.Done()
				ch <- id
			}(i)
		}

		// using a goroutine to close the channel,
		// make sure to run the below goroutine once we have incremented the counter value
		wg.Add(1)
		go func() {
			defer wg.Done()
			wgWorker.Wait()
			close(ch) // sends a signal to stop the range
			//ch <- 100
			// close signal range that no more values be sent and it can stop after receiving remaining values
			// once the channel is closed, we can't send more values to it
		}()
	}()

	//for i := 1; i <= 5; i++ {
	//	fmt.Println(<-ch)
	//}

	// it would run infinitely, channel needs to be closed to stop this range
	// range is a recv call, and receive would block until there is no send
	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()
}
