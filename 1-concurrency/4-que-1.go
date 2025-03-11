package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()

}

func work(workId int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("doing a sub work", workId)
	}()
	fmt.Println("working", workId)
}
