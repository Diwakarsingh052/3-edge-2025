package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// if cancel is called directly without defer , it would cancel the timer immediately
	// without waiting for the timer to finish
	// it should be called in defer for clearing the resources
	defer cancel()
	n, err := SlowFnV2(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SlowFnV2:", n)
}

// context must be the first parameter in the function signature

func SlowFnV2(ctx context.Context) (int, error) {
	time.Sleep(time.Second * 1)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	select {
	// if context is cancelled, then we don't return the success result
	// context cancelled means that user have asked for work cancellation
	case <-ctx.Done():
		fmt.Println("SlowFnV2 work is reversed")
		return 0, ctx.Err()
		// if context is not canceled, then we move on and send the valid result
		//default is always a non-blocking select
	default:
	}
	return 100, nil
}

func selectWorking() {

	// if default is added, the select is never blocking, default is always true case
	select {
	case <-time.After(time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("default case ran, it never block")
		fmt.Println("it creates an unblocking select")
	}
	fmt.Println("done")
}
