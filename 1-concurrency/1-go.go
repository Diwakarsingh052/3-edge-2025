package main

import (
	"fmt"
	"runtime"
)

// concurrency is dealing with a lot of things at once
// parallelism is doing multiple things at once

func main() {
	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(runtime.GOMAXPROCS(1))
}
