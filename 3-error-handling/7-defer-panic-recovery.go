package main

import (
	"fmt"
	"runtime/debug"
)

// panic is a runtime exception
// we need to decide at what leve we need to stop panic
// the level we stop panic must stop or that function must stop that is recovering the panic

// if the caller function doesn't depends on the called function you can stop panic propagation back
// by calling the recovery function in defer in the function that is getting called

// defer guarantees to run // so it would recover the panic if it would happen
func main() {

	DoSomething()
	fmt.Println("end of the main")
}

func DoSomething() {
	// we need to decide where we would recover from the panic
	// the func where we decide to recover the panic needs to stop

	// RecoverPanic would recover the current function from panic, but the function needs to stop
	// it can't continue executing
	defer recoverPanic()
	UpdateSlice(nil)
	fmt.Println("end of DoSomething")
}

func UpdateSlice(s []int) {
	s[0] = 100
	fmt.Println("end of the update slice")
}

func recoverPanic() {

	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.

	// msg would have the actual panic message if that happened
	msg := recover()
	if msg != nil {
		// if this condition is true then panic happened
		fmt.Println(msg)
		fmt.Println(string(debug.Stack()))
	}
}
