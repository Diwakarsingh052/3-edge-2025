package main

import (
	"context"
	"fmt"
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
	ch := make(chan int)

	// putting the timeout value in the context container
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel() // it would cleanup the resources taken up by the context

	go func() {
		x := SlowFn()
		ch <- x
	}()
	select {
	// if received value in time, this case evaluates
	case x := <-ch:
		fmt.Println(x)

		// listen over the done channel and if the time is up this case evaluates
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("cancelling the goroutine")
		return
	}

}

func SlowFn() int {
	time.Sleep(time.Second * 2)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
