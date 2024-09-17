package main

import (
	"log"
	"sync"
)

/*
Example of synchronizing the main goroutine with a child goroutine
such that the main goroutine waits until the child goroutine finishes running
so we get the expected sum value instead of default 0
*/
func main() {
	s1 := 0
	// returns a pointer to a waitGroup so we're not referencing a copy
	wg := &sync.WaitGroup{}
	// adding a counter to indicate the number of concurrent goroutines this waitGroup need to wait for before it can terminate
	wg.Add(1)
	go sum(1, 100, wg, &s1)
	// initiating the waitGroup so main goroutine cannot process any further until Done() is called
	wg.Wait()
	log.Println(s1)
}

func sum(from int, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	/**
	calling Done() method on waitGroup so the main goroutine will stop waiting and continue
	processing further instructions. Telling the main goroutine that this child goroutine is done running
	*/
	wg.Done()
}
