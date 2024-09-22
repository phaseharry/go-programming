// example of using atomic operations to prevent race condition.
package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	s1 := int32(0)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(4)
	sum(1, 25, waitGroup, &s1)
	sum(25, 50, waitGroup, &s1)
	sum(51, 75, waitGroup, &s1)
	sum(76, 100, waitGroup, &s1)
	waitGroup.Wait()

	log.Println(s1)
}

/*
Using go's built in atomic AddInt32 method to avoid race conditions of
getting a stale value and overwriting a necessary sum
*/
func sum(from int, to int, wg *sync.WaitGroup, result *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(result, int32(i))
	}
	wg.Done()
}
