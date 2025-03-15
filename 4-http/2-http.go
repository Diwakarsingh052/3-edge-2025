package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", goroutine)
	http.HandleFunc("/json", sendJson)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func goroutine(w http.ResponseWriter, r *http.Request) {
	// http service can recover from the panic if it happens in the goroutine that
	// is automatically started by go to serve the request
	//panic("this handler func is panicking")

	go func() {
		defer recoverPanic()
		// if panic happens in the new goroutine that you have manually spun up
		// then the program will crash
		// we need manual recover in this case
		panic("panic in the new goroutine")
	}()

}

func sendJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//create a struct with 3 fields
	// add json tag on the struct fields
	// for e.g. FirstName string `json:"first_name"`

	// use json.Marshal to convert the struct to json
	// handle errors
	// to signal errors to end user, use http.Error()
	// send the json using w.Write()

	w.Write([]byte(`{"message": "hello world"}`))
}

func recoverPanic() {

	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.

	// msg would have the actual panic message if that happened
	msg := recover()
	if msg != nil {
		// if this condition is true then panic happened
		fmt.Println(msg)
		//fmt.Println(string(debug.Stack()))
	}
}
