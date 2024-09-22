package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan string) // creates channel with an indefinite size. Will remain open until all messages are consumed.
	go greet(ch)
	/*
		Even though we called greet function. It will be blocked until we actually send a message through to the channel
		so the child goroutine can receive a message. Otherwise, the child goroutine will be blocked
	*/

	ch <- "Hello child Goroutine"
	log.Println(<-ch) // receiving response from child goroutine
	log.Println(<-ch)
}

func greet(ch chan string) {
	msg := <-ch                                 // recieving an incoming message from channel ch
	ch <- fmt.Sprintf("Thanks for \"%s\"", msg) // and sending back 2 messages through channel in response so main goroutine can processing response
	ch <- "Hello main Goroutine"
}
