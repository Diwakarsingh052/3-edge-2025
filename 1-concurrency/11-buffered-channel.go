package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, 5)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("sent", i)
		}

		close(ch)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		defer wg.Done()
		for v := range ch {
			fmt.Println("received ", v)

		}
	}()
	wg.Wait()
}
