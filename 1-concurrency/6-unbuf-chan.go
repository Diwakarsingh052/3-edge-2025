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

	ch <- 1 // send would block until no receivers are present
	go func() {

		x := <-ch
		fmt.Println(x)
	}()

	time.Sleep(10 * time.Millisecond) // use wait groups here

}
