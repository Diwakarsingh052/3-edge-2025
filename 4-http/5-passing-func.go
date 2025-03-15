package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("hello"))
	greet(strings.ToUpper("BOB"))

	//at this time we are calling the function instead of passing it normally
	// when we call a function, its return type should satisfy the caller argument
	// in my case addV2 returns a function which has exact same signature of what operateV3 wants

	operateV3(addV2(), 10, 10)
}
func operateV3(op func(int, int) int, x, y int) {
	sum := op(x, y)
	fmt.Println(sum)
	fmt.Println("ran after calling op variable")

}

func addV2() func(int, int) int {
	f := func(x int, y int) int {
		return x + y
	}
	return f
}

func greet(name string) {

	fmt.Println("hello", name)
}
