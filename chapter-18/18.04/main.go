package main

import "log"

func main() {
	/*
		unbuffered channel (no size included) will have a capacity of 0.
		if we try to put anything in them, the channel won't hold them, instead ,it will block until
		that item can be passed through the channel to a variable. This requires more than 1 goroutine.
	*/
	ch := make(chan string) // creates channel with an indefinite size. Will remain open until all messages are consumed.

	go greet(ch)
	log.Println(<-ch) // main goroutine is waiting on channel ch to receive a value from child goroutine.
}

func greet(ch chan string) {
	ch <- "Hello" // piping "Hello" into channel that's passed in
}
