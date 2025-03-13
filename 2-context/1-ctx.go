package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// cancel a func
// carry a value in request scope

//cancel func

// context would carry timers
// once timer is over context sends a signal to cancel the work that was going on
// now it is on the developer itself to respect context and cancel the work
// if a developer failed to do so, then the work would go on

func main() {
	//Background returns a non-nil, empty Context. we can put timer values in it
	ctx := context.Background()
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	// putting the timeout value in the context container
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel() // it would cleanup the resources taken up by the context

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := SlowFn()
		select {
		case ch <- x: // send
			fmt.Println("value sent over the channel")
		case <-ctx.Done():
			fmt.Println("timeout happened, no longer we can send values")
			fmt.Println("reverse the affect of SlowFunc")
			return
		}

	}()
	func() {

		select {
		// if received value in time, this case evaluates
		case x := <-ch:
			fmt.Println(x)

			// listen over the done channel and if the time is up this case evaluates
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("cancelling the current function")
			fmt.Println("no longer waiting for goroutine to finish")
			return
		}
	}()

	fmt.Println("Doing other important tasks .........")
	wg.Wait()
}

func SlowFn() int {
	time.Sleep(time.Second * 4)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
