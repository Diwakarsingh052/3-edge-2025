package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var err error // error is an interface type
	//errors.New() // it returns a struct that implements the interface
	fmt.Println(strconv.Atoi("abc"))
	fmt.Println(strconv.Atoi("xyz"))
	fmt.Println(strconv.ParseInt("qwerty", 10, 64))
	fmt.Println(strconv.ParseUint("abc", 10, 64))
	fmt.Println(strconv.ParseFloat("efgh", 64))
}
