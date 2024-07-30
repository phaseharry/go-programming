package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	// making a byte slice of size 1 to be a buffer for the contents of the file
	// increasing the slice will increase the amount of memory available to be used to store the contents of the file.
	buffer := make([]byte, 1)
	// infinite loop to iterate through until we've read through all contents of the file

	for {
		// reads the content of file into the buffer slice we created. It returns the number of bytes read into the buffer slice
		bytesRead, err := file.Read(buffer)
		// end of file is considered an error. We want to handle it gracefully by breaking out of the finite loop
		if err == io.EOF {
			break
		}
		// if for some reason, reading in the current set of info leads to error, print it, but continue in read the contents
		// of the file into our buffer slice.
		if err != nil {
			fmt.Println(err)
			continue
		}
		// turning the buffer we've read in to a string so we can see the raw data as oppose to just the byte values.
		if bytesRead > 0 {
			fmt.Print(string(buffer[:bytesRead]))
		}
	}
}
