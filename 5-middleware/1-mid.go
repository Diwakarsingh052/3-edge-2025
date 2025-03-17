package main

import (
	"fmt"
	"net/http"
)

// middleware that exec some pre-processing or the post-processing logic
// req -> mid1->mid-2-> handler->services
func main() {
	http.HandleFunc("/home", Mid(home))
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page")
	w.Write([]byte("home page"))

}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// it is a closure
		// we are using a variable from the outer scope which makes this anonymous func a closure
		fmt.Println("mid layer started")
		fmt.Println("pre processing logic going on")
		next(w, r) // closure
		fmt.Println("post processing logic going on")
		fmt.Println("mid layer finished")
	}

}

// Create one more middleware and add that middleware to /home endpoint
// inside the middleware just print middleware two started
