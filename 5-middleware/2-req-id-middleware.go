package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"net/http"
)

type ctxKey string

const Key ctxKey = "reqIdKey"

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
	ctx := r.Context()

	//fetching the value from the context
	reqId, ok := ctx.Value(Key).(string)
	if !ok {
		reqId = "unknown"
	}
	n := rand.Intn(100)
	if n%2 != 0 {
		log.Println(reqId, "something went wrong", n)
		fmt.Fprintln(w, "can't say hello this time")
		return
	}

	log.Println(reqId, "even number", n)
	fmt.Fprintln(w, "hello user ")

}

func RequestIdMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// put the request id in the context
		reqId := uuid.NewString()

		ctx := r.Context()
		// we would get the updated context, we need to update the request object as well
		ctx = context.WithValue(ctx, Key, reqId)

		fmt.Println("req started with ", reqId)

		//withContext would update the internal context of the request object
		//with the updated context with values or timeouts
		next(w, r.WithContext(ctx))
		fmt.Println("req finished with ", reqId)
	}
}
