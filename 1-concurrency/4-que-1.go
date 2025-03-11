package main

import "fmt"

func main() {

	for i := 1; i <= 4; i++ {
		go work(i)
	}

}

func work(workId int) {
	go func() {
		fmt.Println("doing a sub work", workId)
	}()
	fmt.Println("working", workId)
}
