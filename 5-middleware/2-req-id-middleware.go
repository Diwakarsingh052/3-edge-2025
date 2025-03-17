package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	//fmt.Println(uuid.NewString())
	http.HandleFunc("/hello", RequestIdMiddleware(Hello))
	http.HandleFunc("/", Hello)
	panic(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	//Fprintln directly writes output to the responseWriter
	// it is similar like w.Write()

	// fetch the request id from the context
	n := rand.Intn(100)
	if n%2 != 0 {
		log.Println("something went wrong", n)
		fmt.Fprintln(w, "can't say hello this time")
		return
	}
	fmt.Fprintln(w, "hello user ")

}

func RequestIdMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.NewString()
		// put the request id in the context
		fmt.Println("req started with ", reqId)
		next(w, r)
		fmt.Println("req finished with ", reqId)
	}
}
