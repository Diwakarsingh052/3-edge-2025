package main

import (
	"fmt"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
// channels are only meant to be used in concurrent programming

func main() {

	// unbuffered channel has size of 0
	ch := make(chan int, 0)
	go func() {
		// this would block if no sender is present,
		//and another goroutine from the queue would be picked up
		//which is sender goroutine in this case
		//communication completes, and we get 1 on the screen
		ch <- 1 // send

	}()

	x := <-ch
	fmt.Println(x)

	time.Sleep(10 * time.Millisecond) // use wait groups here

}
