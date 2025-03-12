package main

import (
	"fmt"
	"sync"
)

var x = 1

func main() {
	wg := new(sync.WaitGroup)

	m := new(sync.Mutex)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go UpdateX(i, wg, m)
	}

	wg.Wait()
}

func UpdateX(val int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	// critical section
	// this is the place where we access the shared resource

	// when a goroutine acquires a lock, another goroutine can't access the critical section
	// until the lock is not released

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads
	m.Lock()
	defer m.Unlock()
	x = val
	fmt.Println("UpdateX:", x)
}
