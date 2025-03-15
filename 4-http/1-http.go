package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// if handler is nil
	// it would use default route matcher available from standard lib, also known as DefaultServeMux
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct

	// context object is inside the Request struct
	ctx := r.Context() // this method gives the context from the request object
	time.Sleep(time.Second * 5)

	// checking if client is still connected
	select {
	// this case denotes, client is no longer available
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("user is no longer there to receive the response")
		return
	default:
		// client is still connected
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Hello World!"))
	fmt.Println("sent the hello world response")
}
