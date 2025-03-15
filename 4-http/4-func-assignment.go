package main

import (
	"fmt"
	"strings"
)

/*
	Create a func StringManipulation
	StringManipulation accepts a func with a type func(string)string

	Create two more func
	1. trimSpace  -> strings.TrimSpace
	2. toUpper 		-> strings.ToUpper

	Goal: Pass trimSpace and toUpper to StringManipulation
*/

type stringOps func(str string) string

func main() {

	fmt.Println(StringManipulation(trimSpace, " This needs to be trimmed "))
	fmt.Println(StringManipulation(toUpper, "This needs to be upper"))

}

func StringManipulation(operation stringOps, str string) string {
	return operation(str)
}

func trimSpace(str string) string {
	return strings.TrimSpace(str)
}

func toUpper(str string) string {
	return strings.ToUpper(str)
}
