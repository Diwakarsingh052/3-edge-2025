package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	//wg := &sync.WaitGroup{}

	// waitgroup counter represents number of goroutine we are running
	wg.Add(1) // add 1 to the counter
	go func() {
		/// giving a guarantee that even
		// in case of panic this would decrement the counter
		defer wg.Done()
		fmt.Println("hello")
	}()

	fmt.Println("doing some other stuff in main")
	wg.Wait() // wait until the counter is not 0
	fmt.Println("end of the program")

}
