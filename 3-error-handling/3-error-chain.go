package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func main() {
	_, err := openFileV2("abc.txt")
	if err != nil {
		// checking if custom error is present in the chain or not
		// errors.Is will only work if error was wrapped inside the chain
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("custom error found in  the chain")
			return
		}
		log.Println(err)
		return
	}
}

func openFileV2(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {

		// wrap the error to create a chain of error
		// if we use wrapping then we can look inside the chain using errors.Is
		// use %w for wrapping
		return nil, fmt.Errorf("%w %w", ErrFileNotFound, err)
	}
	return f, nil

}
