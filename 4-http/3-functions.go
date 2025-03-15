package main

import (
	"fmt"
)

// money is a new type, it is not an alias
type money int
type operation func(int, int) int

func main() {
	operate(add, 10, 20)
	operate(sub, 90, 10)
	operateV2(add, 80, 100)

}

func operate(f func(int, int) int, x, y int) {
	fmt.Println(f(x, y))

}

func operateV2(f operation, x, y int) {
	fmt.Println(f(x, y))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
