package main

import (
	"context"
	"fmt"
)

// The provided key must be comparable and should not be of type string
// or any other built-in type to avoid collisions between packages using context.
// Users of WithValue should define their own types for keys.
type ctxKey string

const Key ctxKey = "someKey"

func main() {
	var a any = nil
	// using type assertion to check if value is present inside the empty interface or not
	val, ok := a.(int)
	// if ok is true it means the value is of integer type
	if ok {
		fmt.Printf("The value is %d\n", val)
	}
	ctx := context.Background()

	//WithValue is going to store key value pairs
	ctx = context.WithValue(ctx, Key, true)
	ctx.Value(ctx)
	getValueFromTheContext(ctx)

}

func getValueFromTheContext(ctx context.Context) {
	val := ctx.Value(Key)
	reqId, ok := val.(string)
	if !ok {
		fmt.Println("req Id not found or invalid value")
		return
	}
	fmt.Printf("The value is %s\n", reqId)
}
