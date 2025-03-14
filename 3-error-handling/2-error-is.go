package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := openFile("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	info, _ := f.Stat()
	fmt.Println(info.Name())

}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		//errors.Is can check if an error was wrapped inside the chain or not
		//  if an error was found in the chain, you now know what exactly went wrong
		// you might want to take some actions to fix the issue
		//or maybe just log the additional details
		ok := errors.Is(err, os.ErrNotExist)
		if ok {
			log.Println("attempting to create file")
			f, err := os.Create(filename)
			if err != nil {
				return nil, err
			}
			return f, nil
		}

		// if error happened because of another reason other than ErrNotExist
		return nil, err
	}

	return f, nil

}

// f1 -> f2 -> f3
// [err2,err3]	<-[err2,err3](f2) <-err3(f3)
