package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() // spinning up a goroutine, it doesn't mean that it would
	// run it at the same time
	fmt.Println("end of the program")
	time.Sleep(10 * time.Second) // here is a block
	// os would not waste time here for time.Sleep
	// go scheduler would pick other goroutines waiting to exec
}

func hello() {
	fmt.Println("Hello")
}
