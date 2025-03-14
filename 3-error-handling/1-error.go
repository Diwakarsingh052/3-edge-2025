package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

// variable which would be used to store errors should start with word Err

var ErrNotFound = errors.New("not found")

func main() {
	//var err1 error // default value of err is nil
	// avoid if else branching while checking for error, use return to stop the function
	name, err := FetchRecord(1)
	if err != nil {
		// log.Println + os.Exit() // -> log.Fatal
		//log.Fatal(err) // would quit the app // should only be used when critical parts are failing
		//usually at startup and when there is no point of continuing
		//log.Fatalln(err) //

		log.Println(err)
		return
	}
	fmt.Println(name)

	// we can use same err variable again, because if err happened in the previous case
	// then our program would have not made to this line,
	//a,err := ABC()

}

// error must be the last value to be returned from function

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		err := fmt.Errorf("user %d not found", id)
		log.Println(err)
		//return "", ErrNotFound // whenever error happens, set other values to their defaults
		return "", ErrNotFound
	}
	return name, nil

}
