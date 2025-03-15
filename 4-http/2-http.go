package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/go", goroutine)
	http.HandleFunc("/json", sendJson)
	http.HandleFunc("/jsonv2", sendJsonNewEncoder)
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

	var user struct {
		// fields must be exported of a struct so we can convert it to json
		FirstName string `json:"first_name"` // field tag // we are giving what the name of field should be in json output
		Password  string `json:"-"`          // ignoring the field in JSON output
		Email     string `json:"email"`
	}
	user.FirstName = "abc"
	user.Password = "123"
	user.Email = "abc@gmail.com"

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "problem in parsing json", http.StatusInternalServerError)
		return // don't forget the return
		// otherwise the program would continue to exec
	}

	// send the json using w.Write()
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	//w.Write([]byte(`{"message": "hello world"}`))
}

func sendJsonNewEncoder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user struct {
		// fields must be exported of a struct so we can convert it to json
		FirstName string `json:"first_name"` // field tag // we are giving what the name of field should be in json output
		Password  string `json:"-"`          // ignoring the field in JSON output
		Email     string `json:"email"`
	}
	user.FirstName = "abc"
	user.Password = "123"
	user.Email = "abc@gmail.com"

	// below line would convert struct to json and write the response over responseWriter
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "problem in parsing json", http.StatusInternalServerError)
		return
	}

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
