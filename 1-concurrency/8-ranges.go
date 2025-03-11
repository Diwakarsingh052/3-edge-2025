package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // sends a signal to stop the range
		//ch <- 100
		// close signal range that no more values be sent and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it
	}()

	//for i := 1; i <= 5; i++ {
	//	fmt.Println(<-ch)
	//}

	// it would run infinitely, channel needs to be closed to stop this range
	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()
}
