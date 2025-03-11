package main

import (
	"fmt"
	"sync"
)

func main() {
	add := make(chan int)
	sub := make(chan int)
	mult := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go Add(1, 2, add, wg)
	go Sub(1, 2, sub, wg)
	go Mult(1, 2, mult, wg)
	go CollectResult(add, sub, mult, wg)
	wg.Wait()
}

func Add(a, b int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- a + b

	// send the result to the channel
}

func Sub(a, b int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- a - b
}

func Mult(a, b int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- a * b
}
func CollectResult(add, sub, mult chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-add, <-sub, <-mult)
	//recv all the values
}
