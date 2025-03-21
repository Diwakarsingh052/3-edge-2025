package main

import (
	"fmt"
	"sync"
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

		for i := 1; i <= 10; i++ {

			//we need to block our goroutine before closing the channel
			//because we want to make sure all the work
			// is done and finished
			// wgWorker waitgroup we are using to track number of worker goroutines
			wgWorker.Add(1)
			go func(id int) {
				defer wgWorker.Done()
				ch <- id
			}(i)

		}

		// waiting until all the workers are not finished
		wgWorker.Wait()

		// closing when we are sure all the fanned out goroutines finished processing
		close(ch) // sends a signal to stop the range
		//ch <- 100
		// close signal range that no more values be sent and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it
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
