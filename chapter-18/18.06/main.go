package main

import (
	"log"
	"time"
)

func main() {
	s1 := 0 // final sum variable that all child routines will be summing to
	ch := make(chan int, 100)

	// using 4 goroutines to split up the work of pushing nums to channel
	go push(1, 25, ch)   // goroutine a
	go push(26, 50, ch)  // goroutine b
	go push(51, 75, ch)  // goroutine c
	go push(76, 100, ch) // goroutine d

	// listening to channel events and adding it.
	for c := 0; c < 100; c++ {
		/*
			Since we have 4 concurrent goroutines seperately pushing values to the channel (1 - 25, 26 - 50, 51 - 75, 76 - 100),
			the ordering of the numbers will be uniform to that (this effect is even more noticable since we sleep for a microsecond after every push to channel).
			ie. goroutine a will send 1, then b will send 26, c will send 51, and d will send 76 (ordering might be slightly out of wack)
		*/
		s1 += <-ch // this will block and wait until channel sends back a value
	}
	log.Println(s1)
}

func push(from int, to int, out chan int) {
	for i := from; i <= to; i++ {
		log.Println(i)
		out <- i
		// sleeping for a microsecond so another goroutine can pickup the number we've sent through the channel and process it
		time.Sleep(time.Microsecond)
	}
}
