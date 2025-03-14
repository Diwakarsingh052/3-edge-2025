package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Error type should end with the word Error
// Error types should not be used for any domain data
// Creating error struct could be useful when dealing with dynamic error data

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// Error() method is implemented to implement error interface
func (q *QueryError) Error() string {
	// formatting the output string
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

func SearchSomething(id int) error {
	// we must return the pointer to the struct
	// errors.As need pointer to work
	return &QueryError{
		Func:  "SearchSomething",
		Input: strconv.Itoa(id),
		Err:   errors.New("something went wrong"),
	}
}

func main() {
	err := SearchSomething(10)
	// we can use errors.As to check if a struct is present in the chain or not
	if err != nil {

		var qe *QueryError
		ok := errors.As(err, &qe)
		if ok {
			fmt.Println("struct is present inside the chain")
			fmt.Println("now we can access the individual fields of the struct if we want to")
			fmt.Println(qe)

			// we can access individual fields if needed, or take some specific actions
			fmt.Println(qe.Input)
			return
		}
	}

}
